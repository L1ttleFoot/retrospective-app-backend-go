package boards_repository

import (
	"context"
	"errors"
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *BoardsRepository) GetBoardOwnerId(ctx context.Context, boardId string) (string, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
	SELECT owner_id
	FROM boards
	WHERE id=$1;
	`

	var ownerId string

	err := r.pool.QueryRow(ctx, query, boardId).Scan(&ownerId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", fmt.Errorf("board with id %s not found: %w", boardId, core_errors.ErrNotFound)
		}

		return "", fmt.Errorf("failed to get board owner: %w", err)
	}

	return ownerId, nil
}
