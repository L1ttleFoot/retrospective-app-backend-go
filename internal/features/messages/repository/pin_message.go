package messages_repository

import (
	"context"
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
	"time"
)

func (r *MessagesRepository) PinMessage(ctx context.Context, messageId string) error {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
		UPDATE messages
		SET pinned_at=clock_timestamp()
		WHERE id=$1
	`

	cmdTag, err := r.pool.Exec(ctx, query, messageId)
	if err != nil {
		return fmt.Errorf("failed to pin message: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("message with id %s not found: %w", messageId, core_errors.ErrNotFound)
	}

	return nil
}

func (r *MessagesRepository) UnpinMessage(ctx context.Context, messageId string) error {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
		UPDATE messages
		SET pinned_at=NULL
		WHERE id=$1
	`

	cmdTag, err := r.pool.Exec(ctx, query, messageId)
	if err != nil {
		return fmt.Errorf("failed to pin message: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("message with id %s not found: %w", messageId, core_errors.ErrNotFound)
	}

	return nil
}
