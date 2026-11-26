package sections_transport

import (
	"context"

	"golang-clean-server/internal/core/domain"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	core_middleware "golang-clean-server/internal/core/transport/middleware"
	core_server "golang-clean-server/internal/core/transport/server"
	"net/http"
)

type SectionsHandler struct {
	sectionsService SectionsService

	tokenManager core_service_tokens.TokenManager
}

type SectionsService interface {
	CreateSections(ctx context.Context, sections []domain.Section) ([]domain.Section, error)
	GetMessagesBySectionId(ctx context.Context, sectionId string, userId *string) ([]domain.Message, error)
}

func NewSectionsHandler(
	sectionsService SectionsService,
	tokenManager core_service_tokens.TokenManager,
) *SectionsHandler {
	return &SectionsHandler{
		sectionsService: sectionsService,
		tokenManager:    tokenManager,
	}
}

func (h *SectionsHandler) Routes() []core_server.Route {
	return []core_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/sections",
			Handler: h.CreateSections,
		},
		{
			Method:     http.MethodGet,
			Path:       "/sections/{sectionId}/messages",
			Handler:    h.GetMessagesBySectionId,
			Middleware: []core_middleware.Middleware{core_middleware.Auth(h.tokenManager, true)},
		},
	}
}
