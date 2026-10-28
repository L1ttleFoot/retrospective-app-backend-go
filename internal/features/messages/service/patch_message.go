package messages_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	core_errors "golang-clean-server/internal/core/errors"
)

func (s *MessagesService) PatchMessage(ctx context.Context, messageId string, messagePatch domain.MessagePatch) (domain.Message, error) {

	if err := messagePatch.ValidatePatchMessage(); err != nil {
		return domain.Message{}, fmt.Errorf("validate message patch: %w", err)
	}

	currentMessage, err := s.messagesRepository.GetMessageById(ctx, messageId)
	if err != nil {
		return domain.Message{}, fmt.Errorf("get message by id: %w", err)
	}

	fromSectionID := currentMessage.SectionID

	currentMessage.ApplyPatch(messagePatch)

	if err := currentMessage.ValidateMessage(); err != nil {
		return domain.Message{}, fmt.Errorf("validate patched message: %w", err)
	}

	if currentMessage.SectionID != fromSectionID {
		fromBoardID, err := s.messagesRepository.GetBoardIdBySectionId(ctx, fromSectionID)
		if err != nil {
			return domain.Message{}, fmt.Errorf("get board of source section: %w", err)
		}

		toBoardID, err := s.messagesRepository.GetBoardIdBySectionId(ctx, currentMessage.SectionID)
		if err != nil {
			return domain.Message{}, fmt.Errorf("get board of target section: %w", err)
		}

		if fromBoardID != toBoardID {
			return domain.Message{}, fmt.Errorf("%w: сообщение нельзя переместить в секцию другой доски", core_errors.ErrInvalidArgument)
		}
	}

	patchMessage, err := s.messagesRepository.PatchMessage(ctx, currentMessage)
	if err != nil {
		return domain.Message{}, fmt.Errorf("patch user: %w", err)
	}

	event := messageEvent{SectionID: patchMessage.SectionID}
	if patchMessage.SectionID != fromSectionID {
		event.FromSectionID = &fromSectionID
		s.publishMessageEvent(ctx, "message.moved", event)
	} else {
		s.publishMessageEvent(ctx, "message.updated", event)
	}

	return patchMessage, nil
}
