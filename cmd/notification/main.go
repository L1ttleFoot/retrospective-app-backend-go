package main

import (
	"context"
	"fmt"
	core_logger "golang-clean-server/internal/core/logger"
	core_rabbitmq "golang-clean-server/internal/core/rabbitmq"
	notification_repository "golang-clean-server/internal/features/notification/repository"
	notification_service "golang-clean-server/internal/features/notification/service"
	notification_transport "golang-clean-server/internal/features/notification/transport"

	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
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

	logger.Debug("init rabbitmq connection")
	rmqConn, err := core_rabbitmq.NewRabbitMqConnection(core_rabbitmq.NewConfigMust())
	if err != nil {
		panic(err)
	}
	defer rmqConn.Close()

	consumer := core_rabbitmq.NewEventConsumer(
		rmqConn.Connection,
		logger.Logger,
	)

	notificationRepository := notification_repository.NewNotificationRepository("7613385678:AAHOY9cR4ARTepUypsyA5kaD-HYXE2U1940")
	notificationService := notification_service.NewNotificationService(notificationRepository)
	notificationHandler := notification_transport.NewTelegramHandler(notificationService)

	err = consumer.Consume(
		ctx,
		"messages.created",
		func(body []byte) error {

			notificationHandler.SendNotification(body)

			return nil
		},
	)
	if err != nil {
		logger.Fatal("consumer stopped", zap.Error(err))
	}

}
