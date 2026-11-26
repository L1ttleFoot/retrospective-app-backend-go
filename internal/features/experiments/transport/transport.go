package exp_transport

import (
	"context"
	core_server "golang-clean-server/internal/core/transport/server"
	"net/http"
)

type ExpHandler struct {
	experimentsService ExperimentsService
}

type ExperimentsService interface {
	ExpHttp(ctx context.Context) error
}

func NewExpTransport(
	experimentsService ExperimentsService,
) *ExpHandler {
	return &ExpHandler{
		experimentsService: experimentsService,
	}
}

func (h *ExpHandler) Router() []core_server.Route {
	return []core_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/http",
			Handler: h.ExpHttp,
		},
	}
}
