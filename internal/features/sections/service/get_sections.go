package sections_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
)

func (s *SectionsService) GetSections(ctx context.Context, boardId string) ([]domain.Section, error) {

	domainSections, err := s.sectionsRepository.GetSections(ctx, boardId)
	if err != nil {
		return []domain.Section{}, fmt.Errorf("get sections: %w", err)
	}

	return domainSections, nil
}
