package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_kafka "golang-clean-server/internal/core/kafka"
	core_logger "golang-clean-server/internal/core/logger"
	repository_cache "golang-clean-server/internal/core/repository/cache"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

var timeZone = time.UTC

type boardEvent struct {
	EventID   string `json:"eventId"`
	Type      string `json:"type"`
	BoardID   string `json:"boardId"`
	SectionID string `json:"sectionId"`
	MessageID string `json:"messageId"`
	AuthorID  string `json:"authorId"`
}

func main() {
	time.Local = timeZone

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init logger: ", err)
		os.Exit(1)
	}

	cache, err := repository_cache.NewConnectionCache(ctx, repository_cache.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to connect redis cache", zap.Error(err))
	}
	defer cache.Close()

	kafkaConfig := core_kafka.NewConfigMust()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: kafkaConfig.Brokers,
		GroupID: "moderation",
		Topic:   "board-events",
	})
	defer reader.Close()

	writer := &kafka.Writer{
		Addr:         kafka.TCP(kafkaConfig.Brokers...),
		Balancer:     &kafka.Hash{},
		RequiredAcks: kafka.RequireAll,
	}
	defer writer.Close()

	logger.Info("moderation started")

	for {
		msg, err := reader.FetchMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			logger.Error("fetch message", zap.Error(err))
			continue
		}

		var handleErr error
		for attempt := 1; attempt <= 3; attempt++ {
			if handleErr = handle(ctx, cache, writer, msg); handleErr == nil {
				break
			}

			logger.Error("handle message", zap.Int("attempt", attempt), zap.Error(handleErr))

			select {
			case <-ctx.Done():
				return
			case <-time.After(200 * time.Millisecond * time.Duration(attempt)):
			}
		}

		if handleErr != nil {
			for {
				err := writer.WriteMessages(ctx, kafka.Message{
					Topic:   "board-events.dlq",
					Key:     msg.Key,
					Value:   msg.Value,
					Headers: []kafka.Header{{Key: "error", Value: []byte(handleErr.Error())}},
				})
				if err == nil {
					break
				}
				logger.Error("send to dlq, retry", zap.Error(err))
				select {
				case <-ctx.Done():
					return
				case <-time.After(time.Second):
				}
			}
		}

		if err := reader.CommitMessages(ctx, msg); err != nil {
			logger.Error("commit offset", zap.Error(err))
		}
	}
}

type verdictEvent struct {
	EventID   string `json:"eventId"`
	MessageID string `json:"messageId"`
	BoardID   string `json:"boardId"`
	AuthorID  string `json:"authorId"`
	Verdict   string `json:"verdict"`
}

func handle(ctx context.Context, cache *repository_cache.ConnectionCache, writer *kafka.Writer, msg kafka.Message) error {

	var event boardEvent

	if err := json.Unmarshal(msg.Value, &event); err != nil {
		return fmt.Errorf("unmarshal event: %w", err)
	}

	if event.Type != "message.created" {
		return nil
	}

	processedKey := "processed:" + event.EventID
	n, err := cache.Exists(ctx, processedKey).Result()
	if err != nil {
		return fmt.Errorf("check processed: %w", err)
	}
	if n > 0 {
		return nil
	}

	rateKey := "rate:" + event.AuthorID
	count, err := cache.Incr(ctx, rateKey).Result()
	if err != nil {
		return fmt.Errorf("incr rate: %w", err)
	}
	if count == 1 {
		if err := cache.Expire(ctx, rateKey, time.Minute).Err(); err != nil {
			return fmt.Errorf("expire rate: %w", err)
		}
	}

	verdict := "ok"
	if count > 10 {
		verdict = "rate_limited"
	}

	body, err := json.Marshal(verdictEvent{
		EventID:   event.EventID,
		MessageID: event.MessageID,
		BoardID:   event.BoardID,
		AuthorID:  event.AuthorID,
		Verdict:   verdict,
	})
	if err != nil {
		return fmt.Errorf("marshal verdict: %w", err)
	}

	if err := writer.WriteMessages(ctx, kafka.Message{
		Topic: "messages.moderated",
		Key:   []byte(event.BoardID),
		Value: body,
	}); err != nil {
		return fmt.Errorf("publish verdict: %w", err)
	}

	if err := cache.Set(ctx, processedKey, 1, 24*time.Hour); err != nil {
		return fmt.Errorf("mark processed: %w", err)
	}

	return nil
}
