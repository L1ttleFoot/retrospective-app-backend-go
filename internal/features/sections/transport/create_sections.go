package sections_transport

import (
	"golang-clean-server/internal/core/domain"
	core_logger "golang-clean-server/internal/core/logger"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
	"time"
)

type SectionsItemRequest struct {
	Title   string `json:"title"`
	Color   string `json:"color"`
	BoardID string `json:"boardId"`
}

type SectionsItemResponse struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Color   string `json:"color"`
	BoardID string `json:"boardId"`
}

type CreateSectionsRequest struct {
	Sections []SectionsItemRequest `json:"sections" validate:"required"`
}

type CreateSectionsResponse struct {
	Sections []SectionsItemResponse `json:"sections"`
}

func (h *SectionsHandler) CreateSections(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_response.NewResponseHandler(log, w)

	var request CreateSectionsRequest

	if err := core_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	domainSections := make([]domain.Section, 0, len(request.Sections))

	for _, s := range request.Sections {
		newSection := domain.NewSection("", s.Title, s.Color, s.BoardID, time.Time{})
		domainSections = append(domainSections, newSection)
	}

	domainSections, err := h.sectionsService.CreateSections(ctx, domainSections)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create sections")
		return
	}

	responseSections := make([]SectionsItemResponse, 0, len(domainSections))

	for _, s := range domainSections {
		responseSections = append(responseSections, SectionsItemResponse{ID: s.ID, Title: s.Title, Color: s.Color, BoardID: s.BoardID})
	}

	response := CreateSectionsResponse{
		Sections: responseSections,
	}

	responseHandler.JsonResponse(response, http.StatusCreated)
}
