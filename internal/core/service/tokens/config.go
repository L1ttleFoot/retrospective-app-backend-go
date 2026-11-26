package core_service_tokens

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AccessSecret  string `envconfig:"ACCESS_TOKEN_SECRET" required:"true"`
	RefreshSecret string `envconfig:"REFRESH_TOKEN_SECRET" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()

	if err != nil {
		err = fmt.Errorf("get jwt manager config: %w", err)
		panic(err)
	}

	return config
}
