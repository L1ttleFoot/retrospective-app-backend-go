package templates_transport

import (
	core_logger "golang-clean-server/internal/core/logger"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
	"time"
)

type TemplateItemResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	AuthorID  string    `json:"authorId"`
	CreatedAt time.Time `json:"createdAt"`
}

type GetTemplatesResponse struct {
	Templates []TemplateItemResponse `json:"templates"`
}

func (h *TemplatesHandler) GetTemplates(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_response.NewResponseHandler(log, w)

	claims := core_service_tokens.FromContext(ctx)

	templateDomains, err := h.templatesService.GetTemplates(ctx, claims.UserID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get templates")
		return
	}

	responseTemplates := make([]TemplateItemResponse, 0, len(templateDomains))

	for _, template := range templateDomains {
		responseTemplates = append(responseTemplates, TemplateItemResponse{
			ID:        template.ID,
			Title:     template.Title,
			AuthorID:  template.AuthorID,
			CreatedAt: template.CreatedAt,
		})
	}

	response := GetTemplatesResponse{
		Templates: responseTemplates,
	}

	responseHandler.JsonResponse(response, http.StatusOK)
}
