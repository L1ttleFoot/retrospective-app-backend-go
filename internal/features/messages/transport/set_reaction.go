package messages_transport

import (
	core_logger "golang-clean-server/internal/core/logger"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
)

type SetReactionRequest struct {
	ReactionName string `json:"reaction"`
}

func (h *MessagesHandler) SetReaction(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_response.NewResponseHandler(log, w)

	messageId, err := core_request.GetPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "invalid id format")
		return
	}

	claims := core_service_tokens.FromContext(ctx)

	var request SetReactionRequest

	if err := core_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	err = h.messagesService.SetReaction(ctx, request.ReactionName, messageId, claims.UserID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to set reaction")
		return
	}

	responseHandler.EmptyResponse(http.StatusOK)
}
