package templates_transport

import (
	"context"
	"golang-clean-server/internal/core/domain"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	core_middleware "golang-clean-server/internal/core/transport/middleware"
	core_server "golang-clean-server/internal/core/transport/server"
	"net/http"
)

type TemplatesHandler struct {
	templatesService TemplatesService
	tokenManager     core_service_tokens.TokenManager
}

type TemplatesService interface {
	CreateTemplate(ctx context.Context, template domain.Template) (domain.Template, error)
	GetTemplates(ctx context.Context, ownerId string) ([]domain.Template, error)
	DeleteTemplate(ctx context.Context, templateId string) error
}

func NewTemplatesHandler(
	templatesService TemplatesService,
	tokenManager core_service_tokens.TokenManager,
) *TemplatesHandler {
	return &TemplatesHandler{
		templatesService: templatesService,
		tokenManager:     tokenManager,
	}
}

func (h *TemplatesHandler) Router() []core_server.Route {
	return []core_server.Route{
		{
			Method:     http.MethodGet,
			Path:       "/templates",
			Handler:    h.GetTemplates,
			Middleware: []core_middleware.Middleware{core_middleware.Auth(h.tokenManager, false)},
		},
		{
			Method:     http.MethodPost,
			Path:       "/templates",
			Handler:    h.CreateTemplate,
			Middleware: []core_middleware.Middleware{core_middleware.Auth(h.tokenManager, false)},
		},
		{
			Method:     http.MethodDelete,
			Path:       "/templates/{id}",
			Handler:    h.DeleteTemplate,
			Middleware: []core_middleware.Middleware{core_middleware.Auth(h.tokenManager, false)},
		},
	}
}
