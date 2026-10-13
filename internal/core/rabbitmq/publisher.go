package core_rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-clean-server/internal/core/domain"

	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type EventPublisher struct {
	conn   *amqp091.Connection
	logger *zap.Logger
}

func NewEventPublisher(conn *amqp091.Connection, logger *zap.Logger) *EventPublisher {
	return &EventPublisher{
		conn:   conn,
		logger: logger,
	}
}

func (p *EventPublisher) Publish(ctx context.Context, queueName string, body []byte) error {

	ch, err := p.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}

	err = ch.PublishWithContext(ctx,
		"",
		queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	p.logger.Debug("message published successfully", zap.String("queue", queueName))
	return nil
}

func (p *EventPublisher) PublishMessageCreated(ctx context.Context, queue string, msg domain.Message) error {

	bytes, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal message error: %w", err)
	}

	return p.Publish(ctx, queue, bytes)
}
