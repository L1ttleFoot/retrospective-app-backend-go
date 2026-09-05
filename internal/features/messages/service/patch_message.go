package messages_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
)

func (s *MessagesService) PatchMessage(ctx context.Context, messageId string, messagePatch domain.MessagePatch) (domain.Message, error) {

	if err := messagePatch.ValidatePatchMessage(); err != nil {
		return domain.Message{}, fmt.Errorf("validate message patch: %w", err)
	}

	currentMessage, err := s.messagesRepository.GetMessageById(ctx, messageId)
	if err != nil {
		return domain.Message{}, fmt.Errorf("get message by id: %w", err)
	}

	currentMessage.ApplyPatch(messagePatch)

	if err := currentMessage.ValidateMessage(); err != nil {
		return domain.Message{}, fmt.Errorf("validate patched message: %w", err)
	}

	patchMessage, err := s.messagesRepository.PatchMessage(ctx, currentMessage)
	if err != nil {
		return domain.Message{}, fmt.Errorf("patch user: %w", err)
	}

	return patchMessage, nil
}
