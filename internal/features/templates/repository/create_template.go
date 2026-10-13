package templates_repository

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *TemplatesRepository) CreateTemplate(ctx context.Context, template domain.Template) (domain.Template, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
	INSERT INTO templates (title, author_id)
	VALUES ($1,$2)
	RETURNING id, created_at
	`
	fmt.Println(query)

	row := r.pool.QueryRow(ctx, query, template.Title, template.AuthorID)

	if err := row.Scan(&template.ID, &template.CreatedAt); err != nil {
		return domain.Template{}, fmt.Errorf("template scan error: %w", err)
	}

	batch := &pgx.Batch{}

	query = `
	INSERT INTO sections (title, color, templateId)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	for _, s := range template.Sections {
		batch.Queue(query, s.Title, s.Color, template.ID)
	}

	br := r.pool.SendBatch(ctx, batch)
	defer br.Close()

	return domain.Template{}, nil
}
