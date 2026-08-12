package messages_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
)

func (s *MessagesService) PatchMessage(ctx context.Context, messageId string, messagePatch domain.MessagePatch) (domain.Message, error) {

	if err := messagePatch.ValidateMessage(); err != nil {
		return domain.Message{}, fmt.Errorf("validate message patch: %w", err)
	}

	return domain.Message{}, nil
}
