package core_telegram

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Token  string `envconfig:"TG_TOKEN" required:"true"`
	ChatID int64  `envconfig:"TG_CHAT_ID" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {
	cfg, err := NewConfig()
	if err != nil {
		err := fmt.Errorf("get telegram config: %w", err)
		panic(err)
	}

	return cfg
}
