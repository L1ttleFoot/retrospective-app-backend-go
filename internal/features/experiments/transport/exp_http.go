package exp_transport

import (
	core_logger "golang-clean-server/internal/core/logger"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
)

func (h *ExpHandler) ExpHttp(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_response.NewResponseHandler(log, w)

	err := h.experimentsService.ExpHttp(ctx)
	if err != nil {
		responseHandler.ErrorResponse(err, "exp http error")
		return
	}

	responseHandler.EmptyResponse(http.StatusOK)
}
