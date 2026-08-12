package boards_repository

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	"time"
)

func (r *BoardsRepository) CreateBoard(ctx context.Context, board domain.Board) (domain.Board, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
	INSERT INTO boards (title, owner_id)
	VALUES ($1, $2)
	RETURNING id, created_at;
	`

	row := r.pool.QueryRow(ctx, query, board.Title, board.OwnerID)

	if err := row.Scan(&board.ID, &board.CreatedAt); err != nil {
		return domain.Board{}, fmt.Errorf("board scan error: %w", err)
	}

	boardDomain := domain.NewBoard(board.ID, board.Title, board.OwnerID, board.CreatedAt)

	cacheKey := fmt.Sprintf("board:owner:%s", board.OwnerID)
	_ = r.cache.Delete(ctx, cacheKey)

	return boardDomain, nil
}
