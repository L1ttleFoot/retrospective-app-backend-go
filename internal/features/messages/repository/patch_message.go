package messages_repository

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	"time"
)

func (r *MessagesRepository) PatchMessage(ctx context.Context, message domain.Message) (domain.Message, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
	UPDATE messages
	SET text=$1, section_id=$2
	WHERE id=$3
	RETURNING id, text, author_id, created_at, updated_at, pinned_at, section_id
	`

	row := r.pool.QueryRow(ctx, query, message.Text, message.SectionID, message.ID)

	var messageDomain domain.Message

	if err := row.Scan(
		&messageDomain.ID,
		&messageDomain.Text,
		&messageDomain.AuthorID,
		&messageDomain.CreatedAt,
		&messageDomain.UpdatedAt,
		&messageDomain.PinnedAt,
		&messageDomain.SectionID,
	); err != nil {
		return domain.Message{}, fmt.Errorf("scarn error: %w", err)
	}

	fmt.Printf("message: %+v", message)
	fmt.Printf("messageDomain: %+v", messageDomain)

	return messageDomain, nil
}
