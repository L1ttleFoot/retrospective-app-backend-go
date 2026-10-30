package sections_repository

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	core_errors "golang-clean-server/internal/core/errors"
	"time"
)

func (r *SectionsRepository) GetSections(ctx context.Context, boardId string) ([]domain.Section, error) {

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

	var domainSections []domain.Section

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

	if len(domainSections) == 0 {
		return nil, fmt.Errorf("sections from board with id %s not found: %w", boardId, core_errors.ErrNotFound)
	}

	return domainSections, nil
}
