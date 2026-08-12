package messages_service

import (
	"context"
	"fmt"
)

func (s *MessagesService) DeleteMessage(ctx context.Context, messageId string) error {

	if err := s.messagesRepository.DeleteMessage(ctx, messageId); err != nil {
		return fmt.Errorf("delete message: %w", err)
	}

	return nil
}
