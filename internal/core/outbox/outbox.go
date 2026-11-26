package core_outbox

import (
	"context"
	"fmt"
	core_logger "golang-clean-server/internal/core/logger"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func Run(ctx context.Context, pool *pgxpool.Pool, writer *kafka.Writer, logger *core_logger.Logger) error {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if err := publishBatch(ctx, pool, writer); err != nil {
				logger.Error("publish batch", zap.Error(err))
			}
		}
	}
}

type Event struct {
	ID      int64
	Topic   string
	Key     string
	Payload string
}

func publishBatch(ctx context.Context, pool *pgxpool.Pool, writer *kafka.Writer) error {

	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, `
		SELECT id, topic, key, payload FROM outbox
		WHERE published_at IS NULL ORDER BY id LIMIT 100 FOR UPDATE SKIP LOCKED
	`)
	if err != nil {
		return fmt.Errorf("select outbox: %w", err)
	}

	var events []Event
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Topic, &e.Key, &e.Payload); err != nil {
			return fmt.Errorf("scan outbox: %w", err)
		}

		events = append(events, e)

	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("read outbox: %w", err)
	}
	rows.Close()

	for _, e := range events {

		err := writer.WriteMessages(ctx, kafka.Message{
			Topic: e.Topic,
			Key:   []byte(e.Key),
			Value: []byte(e.Payload),
		})

		if err != nil {
			return fmt.Errorf("publish outbox %d: %w", e.ID, err)
		}
		if _, err := tx.Exec(ctx, "UPDATE outbox SET published_at = now() WHERE id = $1", e.ID); err != nil {
			return fmt.Errorf("mark published %d: %w", e.ID, err)
		}
	}

	return tx.Commit(ctx)
}
