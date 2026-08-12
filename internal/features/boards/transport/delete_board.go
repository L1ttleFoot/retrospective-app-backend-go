package boards_transport

import (
	core_logger "golang-clean-server/internal/core/logger"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
)

func (h *BoardsHandler) DeleteBoard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_response.NewResponseHandler(log, w)

	boardId, err := core_request.GetPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "invalid id format")
		return
	}

	if err := h.boardsService.DeleteBoard(ctx, boardId); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete board")
		return
	}

	responseHandler.EmptyResponse(http.StatusNoContent)
}
