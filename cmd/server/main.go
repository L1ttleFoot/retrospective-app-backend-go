package main

import (
	"context"
	"fmt"
	core_logger "golang-clean-server/internal/core/logger"
	repository_pool "golang-clean-server/internal/core/repository/pool"
	core_service_hasher "golang-clean-server/internal/core/service/hasher"
	"os"

	core_middleware "golang-clean-server/internal/core/transport/middleware"
	core_server "golang-clean-server/internal/core/transport/server"
	user_repository "golang-clean-server/internal/features/users/repository"
	user_service "golang-clean-server/internal/features/users/service"
	users_transport_http "golang-clean-server/internal/features/users/transport/http"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	logger, err := core_logger.NewLogger((core_logger.NewConfigMust()))

	if err != nil {
		fmt.Println("failed to init logger : ", err)
		os.Exit(1)
	}

	logger.Debug("starting app")

	logger.Debug("init postgres connection pool")
	pool, err := repository_pool.NewConnectionPool(ctx, repository_pool.NewConfigMust())

	if err != nil {
		logger.Fatal("failed to connect postgres pool", zap.Error(err))
	}
	defer pool.Close()

	hasher := core_service_hasher.NewHasher(bcrypt.DefaultCost)

	usersRepository := user_repository.NewUsersRepository(pool)
	usersService := user_service.NewUsersService(usersRepository, hasher)
	usersTransport := users_transport_http.NewUsersHandlers(usersService)

	logger.Debug("init http server")
	httpServer := core_server.NewServer(
		core_server.NewConfigMust(),
		logger,
		core_middleware.RequestID(),
		core_middleware.Logger(logger),
		core_middleware.Trace(),
		core_middleware.Panic(),
	)

	apiRouter := core_server.NewApiRouter()

	apiRouter.RegisterRoutes(usersTransport.Routes()...)

	httpServer.RegisterApiRouters(apiRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("http server run error", zap.Error(err))
	}
}
