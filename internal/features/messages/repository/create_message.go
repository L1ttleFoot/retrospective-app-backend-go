package messages_repository

import (
	"context"
	"fmt"
	"golang-clean-server/db"
	"golang-clean-server/internal/core/domain"
	"time"
)

func (r *MessagesRepository) CreateMessage(ctx context.Context, message domain.Message) (domain.Message, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	row, err := r.q.CreateMessage(ctx, db.CreateMessageParams{
		Text:      message.Text,
		AuthorID:  message.AuthorID,
		SectionID: message.SectionID,
	})

	if err != nil {
		return domain.Message{}, fmt.Errorf("create message: %w", err)
	}

	messageDomain := domain.NewMessage(
		row.ID,
		message.Text,
		message.AuthorID,
		row.CreatedAt,
		row.UpdatedAt,
		row.PinnedAt,
		message.SectionID,
	)

	return messageDomain, nil
}

func toDomainMessage(row db.Message) domain.Message {
	return domain.NewMessage(
		row.ID,
		row.Text,
		row.AuthorID,
		row.CreatedAt,
		row.UpdatedAt,
		row.PinnedAt,
		row.SectionID,
	)
}
