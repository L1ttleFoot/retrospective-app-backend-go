package core_init_features

import (
	core_rabbitmq "golang-clean-server/internal/core/rabbitmq"
	repository_pool "golang-clean-server/internal/core/repository/pool"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	core_sse "golang-clean-server/internal/core/transport/SSE"
	messages_repository "golang-clean-server/internal/features/messages/repository"
	messages_service "golang-clean-server/internal/features/messages/service"
	messages_transport "golang-clean-server/internal/features/messages/transport"
)

func InitMessages(
	pool *repository_pool.ConnectionPool,
	tokenManager core_service_tokens.TokenManager,
	rmqPublisher *core_rabbitmq.EventPublisher,
	eventBroker *core_sse.Broker,
) *messages_transport.MessagesHandler {

	messagerRepository := messages_repository.NewMessagesRepository(pool)
	messagesService := messages_service.NewMessagesService(messagerRepository, rmqPublisher, eventBroker)
	messagesTransport := messages_transport.NewMessagesTransport(messagesService, tokenManager)

	return messagesTransport
}
