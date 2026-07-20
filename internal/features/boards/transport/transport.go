package boards_transport

import (
	"context"
	core_server "golang-clean-server/internal/core/transport/server"
	"net/http"
)

type BoardsHandler struct {
	boardsService BoardsService
}

type BoardsService interface {
	CreateBoard(ctx context.Context) (any, error)
	GerBoadr(ctx context.Context, id *string) (any, error)
}

func (h *BoardsHandler) Router() []core_server.Route {
	return []core_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/boards",
			Handler: h.CreateBoard,
		},
		{
			Method:  http.MethodGet,
			Path:    "/boards",
			Handler: h.GetBoardsByOwnerId,
		},
	}
}
