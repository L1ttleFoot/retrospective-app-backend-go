package messages_repository

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	"time"
)

func (r *MessagesRepository) CreateMessage(ctx context.Context, message domain.Message) (domain.Message, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
		INSERT INTO messages (text, author_id, section_id)
		VALUES ($1,$2,$3)
		RETURNING id, created_at, updated_at
	`

	row := r.pool.QueryRow(ctx, query, message.Text, message.AuthorID, message.SectionID)

	if err := row.Scan(&message.ID, &message.CreatedAt, &message.UpdatedAt); err != nil {
		return domain.Message{}, fmt.Errorf("board scan error: %w", err)
	}

	messageDomain := domain.NewMessage(
		message.ID,
		message.Text,
		message.AuthorID,
		message.CreatedAt,
		message.UpdatedAt,
		message.SectionID,
	)

	return messageDomain, nil
}
