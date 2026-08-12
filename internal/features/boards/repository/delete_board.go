package boards_repository

import (
	"context"
	"errors"
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *BoardsRepository) DeleteBoard(ctx context.Context, boardId string) error {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
	DELETE FROM boards
	WHERE id=$1
	RETURNING owner_id;
	`

	var ownerId string

	err := r.pool.QueryRow(ctx, query, boardId).Scan(&ownerId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("board with id %s not found: %w", boardId, core_errors.ErrNotFound)
		}

		return fmt.Errorf("failed to delete board: %w", err)
	}

	cacheKey := fmt.Sprintf("board:owner:%s", ownerId)
	if err := r.cache.Delete(ctx, cacheKey); err != nil {

	}

	return nil
}
