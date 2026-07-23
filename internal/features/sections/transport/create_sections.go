package sections_transport

import (
	core_logger "golang-clean-server/internal/core/logger"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
)

type SectionsItem struct {
	Title   string `json:"title"`
	Color   string `json:"color"`
	BoardId string `json:"boardId"`
}

type CreateSectionsRequest struct {
	Sections []SectionsItem `json:"sections"`
}

func (h *SectionsHandler) CreateSections(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_response.NewResponseHandler(log, w)

	responseHandler.EmptyResponse(http.StatusOK)
}
