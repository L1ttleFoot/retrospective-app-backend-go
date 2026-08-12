package sections_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
)

func (s *SectionsService) CreateSections(ctx context.Context, sections []domain.Section) ([]domain.Section, error) {

	domainSections, err := s.sectionsRepository.CreateSections(ctx, sections)
	if err != nil {
		return []domain.Section{}, fmt.Errorf("create sections %w", err)
	}

	return domainSections, nil
}
