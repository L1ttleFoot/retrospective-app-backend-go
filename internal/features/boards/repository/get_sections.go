package boards_repository

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	"time"
)

func (r *BoardsRepository) GetSections(ctx context.Context, boardId string) ([]domain.Section, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
	SELECT id, title, color, board_id, created_at
	FROM sections
	WHERE board_id = $1
	`

	rows, err := r.pool.Query(ctx, query, boardId)
	if err != nil {
		return []domain.Section{}, fmt.Errorf("pool error: %w", err)
	}
	defer rows.Close()

	domainSections := make([]domain.Section, 0)

	for rows.Next() {
		var s domain.Section

		if err := rows.Scan(&s.ID, &s.Title, &s.Color, &s.BoardID, &s.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan user row: %w", err)
		}

		domainSections = append(domainSections, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	return domainSections, nil
}
