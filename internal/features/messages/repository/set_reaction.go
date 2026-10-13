package messages_repository

import (
	"context"
	"fmt"
	"time"
)

func (r *MessagesRepository) SetReaction(ctx context.Context, reaction, messageId, authorId string) error {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
	WITH deleted AS (
		DELETE FROM messages_reactions
		WHERE reaction=$1 AND message_id=$2 AND author_id=$3
		RETURNING 1
	)
	INSERT INTO messages_reactions (reaction, message_id, author_id)
	SELECT $1,$2,$3
	WHERE NOT EXISTS (SELECT 1 FROM deleted)
	`

	_, err := r.pool.Exec(ctx, query, reaction, messageId, authorId)

	/* query := `
	INSERT INTO reactions (reaction)
	VALUES('thumbsUp'),
	('thumbsDown'),
	('fire')
	`

	_, err := r.pool.Exec(ctx, query) */

	if err != nil {
		return fmt.Errorf("failed to set reaction: %w", err)
	}

	return nil
}
