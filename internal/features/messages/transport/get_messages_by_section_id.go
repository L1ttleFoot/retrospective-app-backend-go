package messages_transport

import (
	core_logger "golang-clean-server/internal/core/logger"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
	"time"
)

type MessagesReactionResponse struct {
	Name       string `json:"name"`
	Char       string `json:"char"`
	Count      int    `json:"count"`
	IsSelected bool   `json:"isSelected"`
}

type MessagesResponse struct {
	ID        string                     `json:"id"`
	Text      string                     `json:"text"`
	AuthorId  string                     `json:"authorId"`
	CreatedAt time.Time                  `json:"createdAt"`
	UpdatedAt time.Time                  `json:"updatedAt"`
	PinnedAt  *time.Time                 `json:"pinnedAt"`
	SectionId string                     `json:"sectionId"`
	Reactions []MessagesReactionResponse `json:"reactions"`
}

type GetMessagesBySectionIdResponse struct {
	Messages []MessagesResponse `json:"messages" validate:"required"`
}

func (h *MessagesHandler) GetMessagesBySectionId(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_response.NewResponseHandler(log, w)

	sectionId, err := core_request.GetPathValue(r, "sectionId")
	if err != nil {
		responseHandler.ErrorResponse(err, "invalid id format")
		return
	}

	var userId *string
	if claims := core_service_tokens.FromContextOptional(ctx); claims != nil {
		userId = &claims.UserID
	}

	messageDomains, err := h.messagesService.GetMessagesBySectionId(ctx, sectionId, userId)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get messages list")
		return
	}

	messagesRespons := make([]MessagesResponse, 0, len(messageDomains))

	for _, v := range messageDomains {

		reactionsResponse := make([]MessagesReactionResponse, 0, len(v.Reactions))

		for _, r := range v.Reactions {
			reactionsResponse = append(reactionsResponse, MessagesReactionResponse{
				Name:       r.Name,
				Char:       r.Char,
				Count:      r.Count,
				IsSelected: r.IsSelected,
			})
		}

		messagesRespons = append(messagesRespons, MessagesResponse{
			ID:        v.ID,
			Text:      v.Text,
			AuthorId:  v.AuthorID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			PinnedAt:  v.PinnedAt,
			SectionId: v.SectionID,
			Reactions: reactionsResponse,
		})
	}

	response := GetMessagesBySectionIdResponse{Messages: messagesRespons}

	responseHandler.JsonResponse(response, http.StatusOK)
}
