package core_rabbitmq

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type ConnectionRabbit struct {
	*amqp091.Connection
}

func NewRabbitMqConnection(config Config) (*ConnectionRabbit, error) {

	connectionString := fmt.Sprintf(
		"amqp://%s:%s@%s:%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
	)

	conn, err := amqp091.Dial(connectionString)
	if err != nil {
		return nil, fmt.Errorf("create rabbitmq conn: %w", err)
	}

	return &ConnectionRabbit{Connection: conn}, nil
}
