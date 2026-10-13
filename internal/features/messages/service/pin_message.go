package messages_service

import (
	"context"
	"fmt"
)

func (s *MessagesService) PinMessage(ctx context.Context, messageId string) error {

	if err := s.messagesRepository.PinMessage(ctx, messageId); err != nil {
		return fmt.Errorf("pin message :%w", err)
	}

	return nil
}

func (s *MessagesService) UnpinMessage(ctx context.Context, messageId string) error {

	if err := s.messagesRepository.UnpinMessage(ctx, messageId); err != nil {
		return fmt.Errorf("unpin message :%w", err)
	}

	return nil
}
