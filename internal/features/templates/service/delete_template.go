package templates_service

import (
	"context"
	"fmt"
)

func (s *TemplatesService) DeleteTemplate(ctx context.Context, templateId string) error {

	if err := s.templateRepository.DeleteTemplate(ctx, templateId); err != nil {
		return fmt.Errorf("delete message: %w", err)
	}

	return nil
}
