package core_rabbitmq

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host     string `envconfig:"HOST"     required:"true"`
	Port     string `envconfig:"PORT"     default:"5672"`
	User     string `envconfig:"USER"     required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
}

func NewConfig() (Config, error) {

	var config Config

	if err := envconfig.Process("RABBITMQ", &config); err != nil {
		return Config{}, fmt.Errorf("process envonfig: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {

	config, err := NewConfig()

	if err != nil {
		err = fmt.Errorf("get rabbitmq connection config: %w", err)
		panic(err)
	}

	return config
}
