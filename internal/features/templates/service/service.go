package templates_service

import (
	"context"
	"golang-clean-server/internal/core/domain"
)

type TemplatesService struct {
	templateRepository TemplateRepository
}

type TemplateRepository interface {
	CreateTemplate(ctx context.Context, template domain.Template) (domain.Template, error)
	GetTemplates(ctx context.Context, ownerId string) ([]domain.Template, error)
	DeleteTemplate(ctx context.Context, templateId string) error
}

func NewTemplateService(templateRepository TemplateRepository) *TemplatesService {
	return &TemplatesService{templateRepository: templateRepository}
}
