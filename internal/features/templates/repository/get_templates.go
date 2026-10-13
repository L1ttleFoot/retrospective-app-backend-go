package templates_repository

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	"time"
)

func (r *TemplatesRepository) GetTemplates(ctx context.Context, ownerId string) ([]domain.Template, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var templateDomains []domain.Template

	query := `
		SELECT id, title, created_at
		FROM templates
		WHERE author_id=$1
		ORDER BY created_at
	`

	rows, err := r.pool.Query(ctx, query, ownerId)
	if err != nil {
		return []domain.Template{}, fmt.Errorf("pool error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var t domain.Template

		if err := rows.Scan(&t.ID, &t.Title, &t.CreatedAt); err != nil {

		}

		templateDomains = append(templateDomains, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	return templateDomains, nil
}
