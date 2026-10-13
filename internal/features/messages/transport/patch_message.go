package messages_transport

import (
	"golang-clean-server/internal/core/domain"
	core_logger "golang-clean-server/internal/core/logger"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	core_types "golang-clean-server/internal/core/transport/types"
	"net/http"
	"time"
)

type PatchMessageResponse struct {
	ID        string     `json:"id"`
	Text      string     `json:"text"`
	AuthorId  string     `json:"authorId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	PinnedAt  *time.Time `json:"pinnedAt"`
	SectionId string     `json:"sectionId"`
}

type PatchMessageRequest struct {
	Text      core_types.Nullable[string] `json:"text"`
	SectionID core_types.Nullable[string] `json:"sectionId"`
}

func (h *MessagesHandler) PatchMessage(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_response.NewResponseHandler(log, w)

	messageId, err := core_request.GetPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "invalid id format")
		return
	}

	var request PatchMessageRequest

	if err := core_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	messagePatch := messagePatchFromRequest(request)

	messageDomain, err := h.messagesService.PatchMessage(ctx, messageId, messagePatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed tp patch message")
		return
	}

	response := PatchMessageResponse{
		ID:        messageDomain.ID,
		Text:      messageDomain.Text,
		AuthorId:  messageDomain.AuthorID,
		CreatedAt: messageDomain.CreatedAt,
		UpdatedAt: messageDomain.UpdatedAt,
		PinnedAt:  messageDomain.PinnedAt,
		SectionId: messageDomain.SectionID,
	}

	responseHandler.JsonResponse(response, http.StatusOK)
}

func messagePatchFromRequest(request PatchMessageRequest) domain.MessagePatch {
	return domain.NewMessagePatch(request.Text.ToDomain(), request.SectionID.ToDomain())
}
