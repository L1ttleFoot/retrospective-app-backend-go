package messages_transport

import (
	"golang-clean-server/internal/core/domain"
	core_logger "golang-clean-server/internal/core/logger"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
	"time"
)

type CreateMessageRequest struct {
	Text      string `json:"text"      validate:"required"`
	SectionId string `json:"sectionId" validate:"required"`
}

type CreateMessageResponse struct {
	ID        string     `json:"id"`
	Text      string     `json:"text"`
	AuthorId  string     `json:"authorId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	PinnedAt  *time.Time `json:"pinnedAt"`
	SectionId string     `json:"sectionId"`
}

func (h *MessagesHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_response.NewResponseHandler(log, w)

	claims := core_service_tokens.FromContext(ctx)

	var request CreateMessageRequest

	if err := core_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	messageDomain, err := h.messagesService.CreateMessage(
		ctx,
		domain.Message{
			Text:      request.Text,
			SectionID: request.SectionId,
			AuthorID:  claims.UserID,
		},
	)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create board")
		return
	}

	response := CreateMessageResponse{
		ID:        messageDomain.ID,
		Text:      messageDomain.Text,
		AuthorId:  messageDomain.AuthorID,
		CreatedAt: messageDomain.CreatedAt,
		UpdatedAt: messageDomain.UpdatedAt,
		PinnedAt:  messageDomain.PinnedAt,
		SectionId: messageDomain.SectionID,
	}

	responseHandler.JsonResponse(response, http.StatusCreated)
}
