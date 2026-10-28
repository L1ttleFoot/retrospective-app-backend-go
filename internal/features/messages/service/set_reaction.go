package messages_service

import (
	"context"
	"fmt"
)

func (s *MessagesService) SetReaction(ctx context.Context, reactionName, messageId, authorId string) error {

	message, err := s.messagesRepository.GetMessageById(ctx, messageId)
	if err != nil {
		return fmt.Errorf("get message by id: %w", err)
	}

	err = s.messagesRepository.SetReaction(ctx, reactionName, messageId, authorId)
	if err != nil {
		return fmt.Errorf("set reaction: %w", err)
	}

	s.publishMessageEvent(ctx, "message.updated", messageEvent{SectionID: message.SectionID})

	return nil
}
