package boards_repository

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	core_logger "golang-clean-server/internal/core/logger"
	"time"

	"go.uber.org/zap"
)

func (r *BoardsRepository) GetBoardsByOwnerId(ctx context.Context, ownerId string) ([]domain.Board, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	log := core_logger.FromContext(ctx)

	cacheKey := fmt.Sprintf("board:owner:%s", ownerId)

	var boardDomains []domain.Board

	if err := r.cache.Get(ctx, cacheKey, &boardDomains); err == nil {
		log.Debug("get data from cache", zap.String("key", cacheKey))

		return boardDomains, nil
	}

	query := `
		SELECT id, title, owner_id, created_at
		FROM boards
		WHERE owner_id = $1;
	`

	rows, err := r.pool.Query(ctx, query, ownerId)
	if err != nil {
		return []domain.Board{}, fmt.Errorf("pool error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var b domain.Board

		if err := rows.Scan(&b.ID, &b.Title, &b.OwnerID, &b.CreatedAt); err != nil {
			return []domain.Board{}, fmt.Errorf("scan board row: %w", err)
		}

		boardDomains = append(boardDomains, b)
	}

	if err := rows.Err(); err != nil {
		return []domain.Board{}, fmt.Errorf("next rows: %w", err)
	}

	_ = r.cache.Set(ctx, cacheKey, boardDomains, 30*time.Second)

	return boardDomains, nil
}
