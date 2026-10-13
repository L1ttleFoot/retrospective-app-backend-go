package templates_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
)

func (s *TemplatesService) GetTemplates(ctx context.Context, ownerId string) ([]domain.Template, error) {

	domainTemplate, err := s.templateRepository.GetTemplates(ctx, ownerId)
	if err != nil {
		return []domain.Template{}, fmt.Errorf("get templates: %w", err)
	}

	return domainTemplate, nil
}
