package messages_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-clean-server/db"
	"golang-clean-server/internal/core/domain"
	"time"

	"github.com/google/uuid"
)

type MessageEvent struct {
	EventId    string    `json:"eventId"`
	Type       string    `json:"type"`
	BoardId    string    `json:"boardId"`
	SectionId  string    `json:"sectionId"`
	MessageId  string    `json:"messageId"`
	AuthorId   string    `json:"authorId"`
	OccurredAt time.Time `json:"occurredAt"`
}

func (r *MessagesRepository) CreateMessageWithOutbox(ctx context.Context, message domain.Message) (domain.Message, error) {

	eventID := uuid.NewString()

	boardId, err := r.GetBoardIdBySectionId(ctx, message.SectionID)
	if err != nil {
		return domain.Message{}, fmt.Errorf("create message: %w", err)
	}

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.Message{}, fmt.Errorf("create message tx: %w", err)
	}
	defer tx.Rollback(ctx)

	queries := r.q.WithTx(tx)

	row, err := queries.CreateMessage(ctx, db.CreateMessageParams{
		Text:      message.Text,
		AuthorID:  message.AuthorID,
		SectionID: message.SectionID,
	})
	if err != nil {
		return domain.Message{}, fmt.Errorf("create message: %w", err)
	}

	payload, err := json.Marshal(MessageEvent{
		EventId:    eventID,
		Type:       "message.created",
		BoardId:    boardId,
		SectionId:  message.SectionID,
		MessageId:  row.ID,
		AuthorId:   message.AuthorID,
		OccurredAt: row.CreatedAt,
	})
	if err != nil {
		return domain.Message{}, fmt.Errorf("marshal message: %w", err)
	}

	query := `
		INSERT INTO outbox (event_id,topic,key,payload)
		VALUES ($1,$2,$3,$4)
	`

	_, err = tx.Exec(ctx, query, eventID, "board-events", boardId, string(payload))
	if err != nil {
		return domain.Message{}, fmt.Errorf("create message: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return domain.Message{}, fmt.Errorf("commit: %w", err)
	}

	messageDomain := domain.NewMessage(
		row.ID,
		message.Text,
		message.AuthorID,
		row.CreatedAt,
		row.UpdatedAt,
		row.PinnedAt,
		message.SectionID,
	)

	return messageDomain, nil
}
