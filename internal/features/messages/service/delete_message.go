package messages_service

import (
	"context"
	"fmt"
)

func (s *MessagesService) DeleteMessage(ctx context.Context, messageId string) error {

	message, err := s.messagesRepository.GetMessageById(ctx, messageId)
	if err != nil {
		return fmt.Errorf("get message by id: %w", err)
	}

	if err := s.messagesRepository.DeleteMessage(ctx, messageId); err != nil {
		return fmt.Errorf("delete message: %w", err)
	}

	s.publishMessageEvent(ctx, "message.deleted", messageEvent{SectionID: message.SectionID})

	return nil
}
