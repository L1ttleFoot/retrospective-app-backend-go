package sections_transport

import (
	"fmt"
	"golang-clean-server/internal/core/domain"
	core_logger "golang-clean-server/internal/core/logger"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
)

type SectionsItem struct {
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
	Sections []SectionsItem `json:"sections"`
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

	domainSections := make([]domain.Section, len(request.Sections))

	for i, s := range request.Sections {
		domainSections[i] = domain.Section{Title: s.Title, Color: s.Color, BoardID: s.BoardID}
	}

	domainSections, err := h.sectionsService.CreateSections(ctx, domainSections)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create sections")
		return
	}

	responseSections := make([]SectionsItemResponse, len(domainSections))

	fmt.Printf("domain sections: %+v", domainSections)

	for i, s := range domainSections {
		responseSections[i] = SectionsItemResponse{ID: s.ID, Title: s.Title, Color: s.Color, BoardID: s.BoardID}
	}

	response := CreateSectionsResponse{
		Sections: responseSections,
	}

	responseHandler.JsonResponse(response, http.StatusCreated)
}
