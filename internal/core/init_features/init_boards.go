package core_init_features

import (
	repository_cache "golang-clean-server/internal/core/repository/cache"
	repository_pool "golang-clean-server/internal/core/repository/pool"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	boards_repository "golang-clean-server/internal/features/boards/repository"
	boards_service "golang-clean-server/internal/features/boards/service"
	boards_transport "golang-clean-server/internal/features/boards/transport"
	"golang-clean-server/internal/features/events"
)

func InitBoards(
	pool *repository_pool.ConnectionPool,
	cache *repository_cache.ConnectionCache,
	tokenManager *core_service_tokens.Tokens,
	eventBroker *events.Broker,
) *boards_transport.BoardsHandler {

	boardsRepository := boards_repository.NewBoardsRepository(pool, cache)
	boardsService := boards_service.NewBoardsService(boardsRepository)
	boardsTransport := boards_transport.NewBoardsHandler(boardsService, tokenManager, eventBroker)

	return boardsTransport
}
