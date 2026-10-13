package messages_transport

import (
	core_logger "golang-clean-server/internal/core/logger"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
	"time"
)

type MessagesResponse struct {
	ID        string     `json:"id"`
	Text      string     `json:"text"`
	AuthorId  string     `json:"authorId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	PinnedAt  *time.Time `json:"pinnedAt"`
	SectionId string     `json:"sectionId"`
	Reactions []int      `json:"reactions"`
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

	messageDomains, err := h.messagesService.GetMessagesBySectionId(ctx, sectionId)
	if err != nil {
		responseHandler.ErrorResponse(err, " failed to get messages list")
		return
	}

	messagesRespons := make([]MessagesResponse, len(messageDomains))

	for i, v := range messageDomains {
		messagesRespons[i] = MessagesResponse{
			ID:        v.ID,
			Text:      v.Text,
			AuthorId:  v.AuthorID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			PinnedAt:  v.PinnedAt,
			SectionId: v.SectionID,
			Reactions: make([]int, 0),
		}
	}

	response := GetMessagesBySectionIdResponse{Messages: messagesRespons}

	responseHandler.JsonResponse(response, http.StatusOK)
}
