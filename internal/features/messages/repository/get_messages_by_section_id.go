package messages_repository

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	"time"
)

func (r *MessagesRepository) GetMessagesBySectionId(ctx context.Context, sectionId string) ([]domain.Message, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var messageDomains []domain.Message

	query := `
	SELECT id, text, author_id, created_at, updated_at, section_id
	FROM messages
	WHERE section_id = $1
	ORDER BY updated_at ASC
	`

	rows, err := r.pool.Query(ctx, query, sectionId)

	if err != nil {
		return []domain.Message{}, fmt.Errorf("pool error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var m domain.Message

		if err := rows.Scan(&m.ID, &m.Text, &m.AuthorID, &m.CreatedAt, &m.UpdatedAt, &m.SectionID); err != nil {
			return nil, fmt.Errorf("scan boadr row: %w", err)
		}

		messageDomains = append(messageDomains, m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	return messageDomains, nil
}
