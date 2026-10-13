package messages_repository

import (
	"context"
	"errors"
	"fmt"
	"golang-clean-server/internal/core/domain"
	core_errors "golang-clean-server/internal/core/errors"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *MessagesRepository) GetMessageById(ctx context.Context, messageId string) (domain.Message, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
	SELECT id, text, author_id, created_at, updated_at, pinned_at, section_id
	FROM messages
	WHERE id = $1
	`

	row := r.pool.QueryRow(ctx, query, messageId)

	var messageDomain domain.Message

	err := row.Scan(
		&messageDomain.ID,
		&messageDomain.Text,
		&messageDomain.AuthorID,
		&messageDomain.CreatedAt,
		&messageDomain.UpdatedAt,
		&messageDomain.PinnedAt,
		&messageDomain.SectionID,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Message{}, fmt.Errorf(
				"message with id=%s: %w",
				messageId,
				core_errors.ErrNotFound,
			)
		}
		return domain.Message{}, fmt.Errorf("scan message row: %w", err)
	}

	return messageDomain, nil
}
