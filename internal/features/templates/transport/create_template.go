package templates_transport

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

type SectionItemRequest struct {
	Title string `json:"title"`
	Color string `json:"color"`
}

type CreateTemplateRequest struct {
	Title    string               `json:"title"`
	Sections []SectionItemRequest `json:"sections"`
}

type SectionItemResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
}

type CreateTemplateResponse struct {
	ID        string                `json:"id"`
	Title     string                `json:"title"`
	Sections  []SectionItemResponse `json:"sections"`
	CreatedAt time.Time             `json:"created_at"`
}

func (h *TemplatesHandler) CreateTemplate(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_response.NewResponseHandler(log, w)

	claims := core_service_tokens.FromContext(ctx)

	var request CreateTemplateRequest

	if err := core_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	domainTemplates := domain.Template{
		Title:    request.Title,
		AuthorID: claims.UserID,
		Sections: make([]domain.TemplateSection, 0, len(request.Sections)),
	}

	for _, v := range request.Sections {
		newTemplateSection := domain.TemplateSection{Title: v.Title, Color: v.Color}

		domainTemplates.Sections = append(domainTemplates.Sections, newTemplateSection)
	}

	fmt.Println(domainTemplates)

	domainTemplates, err := h.templatesService.CreateTemplate(ctx, domainTemplates)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create template")
		return
	}

	response := CreateTemplateResponse{
		ID:        domainTemplates.ID,
		Title:     domainTemplates.Title,
		Sections:  make([]SectionItemResponse, 0, len(domainTemplates.Sections)),
		CreatedAt: domainTemplates.CreatedAt,
	}

	for _, v := range domainTemplates.Sections {

		newSectionTemplate := SectionItemResponse{ID: v.ID, Title: v.Title, Color: v.Color}

		response.Sections = append(response.Sections, newSectionTemplate)
	}

	responseHandler.JsonResponse(response, http.StatusCreated)
}
