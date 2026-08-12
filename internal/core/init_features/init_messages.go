package core_init_features

import (
	repository_pool "golang-clean-server/internal/core/repository/pool"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	messages_repository "golang-clean-server/internal/features/messages/repository"
	messages_service "golang-clean-server/internal/features/messages/service"
	messages_transport "golang-clean-server/internal/features/messages/transport"
)

func InitMessages(
	pool *repository_pool.ConnectionPool,
	tokenManager core_service_tokens.TokenManager,
) *messages_transport.MessagesHandler {

	messagerRepository := messages_repository.NewMessagesRepository(pool)
	messagesService := messages_service.NewMessagesService(messagerRepository)
	messagesTransport := messages_transport.NewMessagesTransport(messagesService, tokenManager)

	return messagesTransport
}
