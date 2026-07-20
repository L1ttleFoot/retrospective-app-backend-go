package boards_transport

import (
	core_logger "golang-clean-server/internal/core/logger"
	core_response "golang-clean-server/internal/core/transport/response"
	core_http_utils "golang-clean-server/internal/core/transport/utils"
	"net/http"
)

func (h *BoardsHandler) GetBoardsByOwnerId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_response.NewResponseHandler(log, w)

	ownerId, err := core_http_utils.GetQueryParamsString(r, "ownerId")
	if err != nil {
		responseHandler.ErrorResponse(err, " failed to get query params")
	}

	boardDomain, err := h.boardsService.GerBoadr(ctx, ownerId)
	if err != nil {
		responseHandler.ErrorResponse(err, " failed to get query params")
	}

	responseHandler.JsonResponse(boardDomain, http.StatusOK)
}
