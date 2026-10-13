package templates_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
)

func (s *TemplatesService) CreateTemplate(ctx context.Context, template domain.Template) (domain.Template, error) {

	if err := template.Validate(); err != nil {
		return domain.Template{}, fmt.Errorf("validate template: %w", err)
	}

	domainTemplate, err := s.templateRepository.CreateTemplate(ctx, template)
	if err != nil {
		return domain.Template{}, fmt.Errorf("create template:%w", err)
	}

	return domainTemplate, nil
}
