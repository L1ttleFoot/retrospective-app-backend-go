package sections_repository

import (
	"context"
	"encoding/json"
	"fmt"
	db "golang-clean-server/db"
	"golang-clean-server/internal/core/domain"
	"time"
)

func (r *SectionsRepository) GetMessagesBySectionId(ctx context.Context, sectionId string, userId *string) ([]domain.Message, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	rows, err := r.q.GetMessagesBySectionID(ctx, db.GetMessagesBySectionIDParams{
		UserID:    userId,
		SectionID: sectionId,
	})
	if err != nil {
		return []domain.Message{}, fmt.Errorf("get messages by section id: %w", err)
	}

	var messageDomains []domain.Message

	for _, row := range rows {
		var reactions []domain.UserReaction
		if err := json.Unmarshal(row.Reactions, &reactions); err != nil {
			return nil, fmt.Errorf("unmarshal reactions: %w", err)
		}

		m := domain.NewMessage(row.ID, row.Text, row.AuthorID, row.CreatedAt, row.UpdatedAt, row.PinnedAt, row.SectionID)
		m.Reactions = reactions
		messageDomains = append(messageDomains, m)
	}

	return messageDomains, nil
}
