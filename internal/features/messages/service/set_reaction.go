package messages_service

import (
	"context"
	"fmt"
)

func (s *MessagesService) SetReaction(ctx context.Context, reaction, messageId, authorId string) error {

	err := s.messagesRepository.SetReaction(ctx, reaction, messageId, authorId)
	if err != nil {
		return fmt.Errorf("set reaction: %w", err)
	}

	return nil
}
