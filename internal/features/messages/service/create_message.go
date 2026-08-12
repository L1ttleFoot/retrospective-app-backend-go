package messages_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
)

func (s *MessagesService) CreateMessage(ctx context.Context, message domain.Message) (domain.Message, error) {

	message, err := s.messagesRepository.CreateMessage(ctx, message)
	if err != nil {
		return domain.Message{}, fmt.Errorf("create board: %w", err)
	}

	return message, nil
}
