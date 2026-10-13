package sections_repository

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *SectionsRepository) CreateSections(ctx context.Context, sections []domain.Section) ([]domain.Section, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	batch := &pgx.Batch{}

	query := `
	INSERT INTO sections (title, color, board_id)
	VALUES ($1, $2, $3)
	RETURNING id, created_at
	`

	for _, s := range sections {
		batch.Queue(query, s.Title, s.Color, s.BoardID)
	}

	br := r.pool.SendBatch(ctx, batch)
	defer br.Close()

	domainSections := make([]domain.Section, 0, len(sections))

	for i, v := range sections {

		var s domain.Section

		err := br.QueryRow().Scan(&s.ID, &s.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("scany error at index %d: %w", i, err)
		}

		domainSections = append(domainSections, domain.NewSection(s.ID, v.Title, v.Color, v.BoardID, s.CreatedAt))
	}

	return domainSections, nil
}
