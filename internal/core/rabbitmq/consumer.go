package core_rabbitmq

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type EventConsumer struct {
	conn   *amqp091.Connection
	logger *zap.Logger
}

func NewEventConsumer(conn *amqp091.Connection, logger *zap.Logger) *EventConsumer {
	return &EventConsumer{
		conn:   conn,
		logger: logger,
	}
}

func (c *EventConsumer) Consume(ctx context.Context, queueName string, handler func([]byte) error) error {

	ch, err := c.conn.Channel()
	if err != nil {
		return fmt.Errorf("open channel: %w", err)
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
		return fmt.Errorf("declare queue: %w", err)
	}

	msgs, err := ch.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("consume: %w", err)
	}

	c.logger.Info(
		"rabbitmq consumer started",
		zap.String("queue", queueName),
	)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case msg, ok := <-msgs:

			if !ok {
				return fmt.Errorf("rabbitmq messages channel closed")
			}

			if err := handler(msg.Body); err != nil {
				c.logger.Error(
					"failed to process message",
					zap.Error(err),
				)

				if err := msg.Nack(false, true); err != nil {
					c.logger.Error(
						"failed to nack message",
						zap.Error(err),
					)
				}

				continue
			}

			if err := msg.Ack(false); err != nil {
				c.logger.Error(
					"failed to ack message",
					zap.Error(err),
				)
			}

		}
	}
}
