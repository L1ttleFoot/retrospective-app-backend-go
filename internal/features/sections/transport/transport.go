package sections_transport

import (
	"context"

	"golang-clean-server/internal/core/domain"
	core_server "golang-clean-server/internal/core/transport/server"
	"net/http"
)

type SectionsHandler struct {
	sectionsService SectionsService
}

type SectionsService interface {
	CreateSections(ctx context.Context) ([]domain.Section, error)
}

func NewSectionsService(sectionsService SectionsService) *SectionsHandler {
	return &SectionsHandler{sectionsService: sectionsService}
}

func (h *SectionsHandler) Routes() []core_server.Route {
	return []core_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/sections",
			Handler: h.CreateSections,
		},
	}
}
