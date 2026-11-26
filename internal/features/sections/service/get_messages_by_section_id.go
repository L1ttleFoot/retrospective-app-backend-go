package sections_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
)

func (s *SectionsService) GetMessagesBySectionId(ctx context.Context, sectionId string, userId *string) ([]domain.Message, error) {

	messages, err := s.sectionsRepository.GetMessagesBySectionId(ctx, sectionId, userId)
	if err != nil {
		return []domain.Message{}, fmt.Errorf("get boards: %w", err)
	}

	return messages, nil
}
