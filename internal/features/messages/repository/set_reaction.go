package messages_repository

import (
	"context"
	"fmt"
	"time"
)

func (r *MessagesRepository) SetReaction(ctx context.Context, reactionName, messageId, authorId string) error {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
		WITH
		target_reaction AS (
			SELECT id FROM reactions WHERE name = $1
		),
		deleted AS (
			DELETE FROM messages_reactions
			WHERE reaction_id = (SELECT id FROM target_reaction)
			AND message_id = $2
			AND author_id = $3
			RETURNING 1
		)
		INSERT INTO messages_reactions (reaction_id, message_id, author_id)
		SELECT (SELECT id FROM target_reaction), $2, $3
		WHERE EXISTS (SELECT 1 FROM target_reaction)
		AND NOT EXISTS (SELECT 1 FROM deleted);
	`

	_, err := r.pool.Exec(ctx, query, reactionName, messageId, authorId)

	if err != nil {
		return fmt.Errorf("failed to set reaction: %w", err)
	}

	return nil
}
