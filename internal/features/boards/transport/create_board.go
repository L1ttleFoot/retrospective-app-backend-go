package boards_transport

import (
	"fmt"
	"golang-clean-server/internal/core/domain"
	core_logger "golang-clean-server/internal/core/logger"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
	"time"
)

type CreateBoardRequest struct {
	Title   string `json:"title" validate:"required"`
	OwnerID string `json:"ownerId"`
}

type CreateBoardResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title" `
	OwnerID   string    `json:"ownerId"`
	CreatedAt time.Time `json:"createdAt"`
}

func (h *BoardsHandler) CreateBoard(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	claims := core_service_tokens.FromContext(ctx)

	fmt.Println(claims)

	var request CreateBoardRequest

	responseHandler := core_response.NewResponseHandler(log, w)

	if err := core_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}
	//OwnerId: claims.OwnerID
	boardDomain, err := h.boardsService.CreateBoard(ctx, domain.Board{Title: request.Title, OwnerId: request.OwnerID})

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create board")
		return
	}

	response := CreateBoardResponse{
		ID:        boardDomain.ID,
		Title:     boardDomain.Title,
		OwnerID:   request.OwnerID,
		CreatedAt: boardDomain.CreatedAt,
	}

	responseHandler.JsonResponse(response, http.StatusCreated)
}
