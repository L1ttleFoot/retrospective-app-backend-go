package main

import (
	"context"
	"fmt"
	core_init_features "golang-clean-server/internal/core/init_features"
	core_logger "golang-clean-server/internal/core/logger"
	core_rabbitmq "golang-clean-server/internal/core/rabbitmq"
	repository_cache "golang-clean-server/internal/core/repository/cache"
	repository_pool "golang-clean-server/internal/core/repository/pool"
	core_service_hasher "golang-clean-server/internal/core/service/hasher"
	core_service_token "golang-clean-server/internal/core/service/tokens"
	"golang-clean-server/internal/features/healthcheck"
	prometheus_transport "golang-clean-server/internal/features/prometheus/transport"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	core_sse "golang-clean-server/internal/core/transport/SSE"
	core_middleware "golang-clean-server/internal/core/transport/middleware"
	core_server "golang-clean-server/internal/core/transport/server"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var timeZone = time.UTC

func main() {

	time.Local = timeZone

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())

	if err != nil {
		fmt.Println("failed to init logger: ", err)
		os.Exit(1)
	}

	logger.Debug("app time zone", zap.Any("zone", timeZone))

	logger.Debug("starting app")

	logger.Debug("init postgres connection pool")
	pool, err := repository_pool.NewConnectionPool(ctx, repository_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to connect postgres pool", zap.Error(err))
	}
	defer pool.Close()

	cache, err := repository_cache.NewConnectionCache(ctx, repository_cache.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to connect redis cache", zap.Error(err))
	}
	defer cache.Close()

	hasher := core_service_hasher.NewHasher(bcrypt.DefaultCost)
	tokenManager := core_service_token.NewTokenManager(core_service_token.NewConfigMust())

	logger.Debug("init sse broker listner")
	eventBroker := core_sse.NewBroker()
	go eventBroker.Listen()

	logger.Debug("init rabbitmq connection")
	rmqConn, err := core_rabbitmq.NewRabbitMqConnection(core_rabbitmq.NewConfigMust())
	if err != nil {
		panic(err)
	}
	defer rmqConn.Close()

	rmqPublisher := core_rabbitmq.NewEventPublisher(rmqConn.Connection, logger.Logger)

	usersTransport := core_init_features.InitUsers(pool, hasher, tokenManager)
	boardsTransport := core_init_features.InitBoards(pool, cache, tokenManager)
	sectionsTransport := core_init_features.InitSections(pool)
	messagesTransport := core_init_features.InitMessages(pool, tokenManager, rmqPublisher, eventBroker)
	templatesTransport := core_init_features.InitTemplates(pool, tokenManager)

	prometheusTransport := prometheus_transport.NewPrometheusHandler()

	logger.Debug("init http server")
	httpServer := core_server.NewServer(
		core_server.NewConfigMust(),
		logger,
		core_middleware.RequestID(),
		core_middleware.Logger(logger),
		core_middleware.Trace(),
		core_middleware.Panic(),
		core_middleware.Metrics(),
	)

	httpServer.RegisterGroups(
		core_server.RouteGroup{
			Prefix: "/api",
			Routes: core_server.RoutesMerge(
				usersTransport.Routes(),
				boardsTransport.Router(),
				sectionsTransport.Routes(),
				messagesTransport.Router(),
				templatesTransport.Router(),
				eventBroker.Routes(),
				prometheusTransport.Routes(),
			),
		},
		core_server.RouteGroup{
			Prefix: "",
			Routes: healthcheck.Routes(),
		},
	)

	go func() {
		logger.Info("pprof listening", zap.String("addr", "0.0.0.0:6060"))
		if err := http.ListenAndServe("0.0.0.0:6060", nil); err != nil {
			logger.Error("pprof server error", zap.Error(err))
		}
	}()

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("http server run error", zap.Error(err))
	}
}
