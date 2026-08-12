package domain

import (
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
	"time"
)

type Message struct {
	ID        string
	Text      string
	AuthorID  string
	CreatedAt time.Time
	UpdatedAt time.Time
	SectionID string
}

func NewMessage(
	id string,
	text string,
	authorId string,
	createdAt time.Time,
	updatedAt time.Time,
	sectionId string,
) Message {
	return Message{
		ID:        id,
		Text:      text,
		AuthorID:  authorId,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		SectionID: sectionId,
	}
}

type MessagePatch struct {
	Text      Nullable[string]
	SectionID Nullable[string]
}

func NewMessagePatch(text, sectionId Nullable[string]) MessagePatch {
	return MessagePatch{Text: text, SectionID: sectionId}
}

func (p *MessagePatch) ValidateMessage() error {

	if p.Text.Set && p.Text.Value == nil {
		return fmt.Errorf("text is required: %w", core_errors.ErrInvalideArgument)
	}

	if p.SectionID.Set && p.SectionID.Value == nil {
		return fmt.Errorf("sectionId is required: %w", core_errors.ErrInvalideArgument)
	}

	return nil
}
