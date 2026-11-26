package reports_transport

import (
	"context"
	"golang-clean-server/internal/core/domain"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	core_middleware "golang-clean-server/internal/core/transport/middleware"
	core_server "golang-clean-server/internal/core/transport/server"
	"net/http"
)

type ReportsHandler struct {
	reportsService ReportsService
	tokenManager   core_service_tokens.TokenManager
}

type ReportsService interface {
	GetReport(ctx context.Context, boardId string, ownerId string) (domain.BoardReport, error)
}

func NewReportsHandler(
	reportsService ReportsService,
	tokenManager core_service_tokens.TokenManager,
) *ReportsHandler {
	return &ReportsHandler{
		reportsService: reportsService,
		tokenManager:   tokenManager,
	}
}

func (h *ReportsHandler) Router() []core_server.Route {
	return []core_server.Route{
		{
			Method:     http.MethodGet,
			Path:       "/reports/boards/{boardId}",
			Handler:    h.GetReport,
			Middleware: []core_middleware.Middleware{core_middleware.Auth(h.tokenManager, false)},
		},
	}
}
