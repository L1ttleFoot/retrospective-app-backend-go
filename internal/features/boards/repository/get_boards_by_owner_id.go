package boards_repository

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	"time"
)

func (r *BoardsRepository) GetBoardsByOwnerId(ctx context.Context, ownerId string) ([]domain.Board, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
		SELECT id, title, owner_id, created_at
		FROM boards
		WHERE owner_id = $1;
	`

	rows, err := r.pool.Query(ctx, query, ownerId)

	var boardDomains []domain.Board

	if err != nil {
		return []domain.Board{}, fmt.Errorf("pool error : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var b domain.Board

		if err := rows.Scan(&b.ID, &b.Title, &b.OwnerId, &b.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan boadr row: %w", err)
		}

		boardDomains = append(boardDomains, b)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows : %w", err)
	}

	return boardDomains, nil
}
