package boards_transport

import (
	core_logger "golang-clean-server/internal/core/logger"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
)

type CreateBoardRequest struct {
	Title   string `json:"title"`
	OwnerID string `json:"ownerId"`
}

type CreateBoardResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	OwnerID   string `json:"ownerId"`
	CreatedAt string `json:"createdAt"`
}

func (h *BoardsHandler) CreateBoard(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	var request CreateBoardRequest

	responseHandler := core_response.NewResponseHandler(log, w)

	if err := core_request.DecodeAndValidate(r, request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
	}

	boardDomain, err := h.boardsService.CreateBoard(ctx)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	responseHandler.JsonResponse(boardDomain, http.StatusCreated)
}
