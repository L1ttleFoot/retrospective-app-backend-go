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

func (m *Message) ValidateMessage() error {

	if m.Text == "" {
		return fmt.Errorf("%w: message text cannot be empty", core_errors.ErrInvalideArgument)
	}

	return nil
}

type MessagePatch struct {
	Text      Nullable[string]
	SectionID Nullable[string]
}

func NewMessagePatch(text, sectionId Nullable[string]) MessagePatch {
	return MessagePatch{Text: text, SectionID: sectionId}
}

func (p *MessagePatch) ValidatePatchMessage() error {

	if p.Text.Set && (p.Text.Value == nil || *p.Text.Value == "") {
		return fmt.Errorf("text is required: %w", core_errors.ErrInvalideArgument)
	}

	if p.SectionID.Set && (p.SectionID.Value == nil || *p.SectionID.Value == "") {
		return fmt.Errorf("sectionId is required: %w", core_errors.ErrInvalideArgument)
	}

	return nil
}

func (m *Message) ApplyPatch(patch MessagePatch) {
	if patch.Text.Set {
		m.Text = *patch.Text.Value
	}
	if patch.SectionID.Set {
		m.SectionID = *patch.SectionID.Value
	}
}
