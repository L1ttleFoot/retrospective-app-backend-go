package boards_repository

import (
	"context"
	"errors"
	"fmt"
	"golang-clean-server/internal/core/domain"
	core_errors "golang-clean-server/internal/core/errors"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *BoardsRepository) GetBoardById(ctx context.Context, boardId string) (domain.Board, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	row, err := r.q.GetBoardById(ctx, boardId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Board{}, fmt.Errorf("board with id %s not found: %w", boardId, core_errors.ErrNotFound)
		}

		return domain.Board{}, fmt.Errorf("get board: %w", err)
	}

	board := domain.NewBoard(
		row.ID,
		row.Title,
		row.OwnerID,
		row.CreatedAt,
	)

	return board, nil
}
