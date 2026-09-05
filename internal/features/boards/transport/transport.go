package boards_transport

import (
	"context"
	"golang-clean-server/internal/core/domain"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	core_sse "golang-clean-server/internal/core/transport/SSE"
	core_middleware "golang-clean-server/internal/core/transport/middleware"
	core_server "golang-clean-server/internal/core/transport/server"

	"net/http"
)

type BoardsHandler struct {
	boardsService BoardsService
	tokenManager  core_service_tokens.TokenManager
	eventBroker   *core_sse.Broker
}

type BoardsService interface {
	CreateBoard(ctx context.Context, board domain.Board) (domain.Board, error)
	GetBoardsByOwnerId(ctx context.Context, ownerId string) ([]domain.Board, error)
	DeleteBoard(ctx context.Context, boardId string) error
}

func NewBoardsHandler(
	boardsService BoardsService,
	tokenManager core_service_tokens.TokenManager,
	eventBroker *core_sse.Broker,
) *BoardsHandler {
	return &BoardsHandler{
		boardsService: boardsService,
		tokenManager:  tokenManager,
		eventBroker:   eventBroker,
	}
}

func (h *BoardsHandler) Router() []core_server.Route {
	return []core_server.Route{
		{
			Method:     http.MethodPost,
			Path:       "/boards",
			Handler:    h.CreateBoard,
			Middleware: []core_middleware.Middleware{core_middleware.Auth(h.tokenManager)},
		},
		{
			Method:     http.MethodGet,
			Path:       "/boards",
			Handler:    h.GetBoardsByOwnerId,
			Middleware: []core_middleware.Middleware{core_middleware.Auth(h.tokenManager)},
		},
		{
			Method:     http.MethodDelete,
			Path:       "/boards/{id}",
			Handler:    h.DeleteBoard,
			Middleware: []core_middleware.Middleware{core_middleware.Auth(h.tokenManager)},
		},
	}
}
