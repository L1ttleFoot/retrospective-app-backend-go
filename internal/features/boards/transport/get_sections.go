package boards_transport

import (
	core_logger "golang-clean-server/internal/core/logger"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
)

type SectionsItemResponse struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Color   string `json:"color"`
	BoardID string `json:"boardId"`
}

type GetSectionsResponse struct {
	BoardTitle string                 `json:"boardTitle"`
	Sections   []SectionsItemResponse `json:"sections"`
}

func (h *BoardsHandler) GetSections(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_response.NewResponseHandler(log, w)

	id, err := core_request.GetPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "invalid id format")
		return
	}

	boardTitle, sectionDomains, err := h.boardsService.GetSections(ctx, id)
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
		BoardTitle: boardTitle,
		Sections:   responseSections,
	}

	responseHandler.JsonResponse(response, http.StatusOK)
}
