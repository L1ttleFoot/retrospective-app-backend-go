package boards_repository

import (
	"context"
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
	"time"
)

func (r *BoardsRepository) DeleteBoard(ctx context.Context, boardId string) error {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
	DELETE FROM boards
	WHERE id=$1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, boardId)

	if err != nil {
		return fmt.Errorf("failed to delete board : %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("board with id %s not found : %w", boardId, core_errors.ErrNotFound)
	}

	return nil
}
