package messages_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
)

func (s *MessagesService) GetMessagesBySectionId(ctx context.Context, sectionId string, userId *string) ([]domain.Message, error) {

	messages, err := s.messagesRepository.GetMessagesBySectionId(ctx, sectionId, userId)
	if err != nil {
		return []domain.Message{}, fmt.Errorf("get boards: %w", err)
	}

	return messages, nil
}
