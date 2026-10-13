package boards_transport

import (
	"encoding/json"
	"golang-clean-server/internal/core/domain"
	core_logger "golang-clean-server/internal/core/logger"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type CreateBoardRequest struct {
	Title string `json:"title" validate:"required"`
}

type CreateBoardResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	OwnerID   string    `json:"ownerId"`
	CreatedAt time.Time `json:"createdAt"`
}

func (h *BoardsHandler) CreateBoard(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	claims := core_service_tokens.FromContext(ctx)

	var request CreateBoardRequest

	responseHandler := core_response.NewResponseHandler(log, w)

	if err := core_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	boardDomain, err := h.boardsService.CreateBoard(ctx, domain.Board{Title: request.Title, OwnerID: claims.UserID})
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create board")
		return
	}

	response := CreateBoardResponse{
		ID:        boardDomain.ID,
		Title:     boardDomain.Title,
		OwnerID:   boardDomain.OwnerID,
		CreatedAt: boardDomain.CreatedAt,
	}

	if jsonBytes, err := json.Marshal(response); err == nil {
		h.eventBroker.Publish(jsonBytes)
		log.Debug("Событие о создании доски отправлено в Pub/Sub")
	} else {
		log.Error("Не удалось сериализовать событие для SSE", zap.Error(err))
	}

	responseHandler.JsonResponse(response, http.StatusCreated)
}
