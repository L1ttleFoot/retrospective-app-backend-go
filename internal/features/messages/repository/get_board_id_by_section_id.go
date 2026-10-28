package messages_repository

import (
	"context"
	"fmt"
	"time"
)

func (r *MessagesRepository) GetBoardIdBySectionId(ctx context.Context, sectionId string) (string, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
	SELECT board_id
	FROM sections
	WHERE id = $1
	`

	var boardId string

	if err := r.pool.QueryRow(ctx, query, sectionId).Scan(&boardId); err != nil {
		return "", fmt.Errorf("get board id by section id: %w", err)
	}

	return boardId, nil
}
