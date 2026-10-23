package domain

import (
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
	"time"
)

type UserReaction struct {
	Name       string
	Char       string
	Count      int
	IsSelected bool
}

type Message struct {
	ID        string
	Text      string
	AuthorID  string
	CreatedAt time.Time
	UpdatedAt time.Time
	PinnedAt  *time.Time
	SectionID string
	Reactions []UserReaction
}

func NewMessage(
	id string,
	text string,
	authorId string,
	createdAt time.Time,
	updatedAt time.Time,
	pinned_at *time.Time,
	sectionId string,
) Message {
	return Message{
		ID:        id,
		Text:      text,
		AuthorID:  authorId,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		PinnedAt:  pinned_at,
		SectionID: sectionId,
	}
}

func (m *Message) ValidateMessage() error {

	if m.Text == "" {
		return fmt.Errorf("%w: message text cannot be empty", core_errors.ErrInvalidArgument)
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
		return fmt.Errorf("text is required: %w", core_errors.ErrInvalidArgument)
	}

	if p.SectionID.Set && (p.SectionID.Value == nil || *p.SectionID.Value == "") {
		return fmt.Errorf("sectionId is required: %w", core_errors.ErrInvalidArgument)
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
