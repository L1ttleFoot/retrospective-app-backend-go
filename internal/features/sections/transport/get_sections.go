package sections_transport

import (
	core_logger "golang-clean-server/internal/core/logger"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
)

type GetSectionsResponse struct {
	Sections []SectionsItemResponse `json:"sections"`
}

func (h *SectionsHandler) GetSections(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_response.NewResponseHandler(log, w)

	id, err := core_request.GetPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "invalid id format")
		return
	}

	sectionDomains, err := h.sectionsService.GetSections(ctx, id)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get sections")
		return
	}

	responseSections := make([]SectionsItemResponse, len(sectionDomains))

	for i, s := range sectionDomains {
		responseSections[i] = SectionsItemResponse{
			ID:      s.ID,
			Title:   s.Title,
			Color:   s.Color,
			BoardID: s.BoardID,
		}
	}

	response := GetSectionsResponse{
		Sections: responseSections,
	}

	responseHandler.JsonResponse(response, http.StatusOK)
}
