package boards_transport

import (
	core_logger "golang-clean-server/internal/core/logger"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type BoardRespons struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	OwnerID   string    `json:"ownerId"`
	CreatedAt time.Time `json:"createdAt"`
}

type GetBoardsByOwnerIdResponse struct {
	Boards []BoardRespons `json:"boards"`
}

func (h *BoardsHandler) GetBoardsByOwnerId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	claims := core_service_tokens.FromContext(ctx)

	log.Debug("claims from context", zap.String("UserID", claims.UserID))

	responseHandler := core_response.NewResponseHandler(log, w)

	boardDomain, err := h.boardsService.GetBoardsByOwnerId(ctx, claims.UserID)
	if err != nil {
		responseHandler.ErrorResponse(err, " failed to get boards list")
	}

	boardsRespons := make([]BoardRespons, len(boardDomain))

	for i, v := range boardDomain {
		boardsRespons[i] = BoardRespons{
			ID:        v.ID,
			Title:     v.Title,
			OwnerID:   v.OwnerID,
			CreatedAt: v.CreatedAt,
		}
	}

	response := GetBoardsByOwnerIdResponse{Boards: boardsRespons}

	responseHandler.JsonResponse(response, http.StatusOK)
}
