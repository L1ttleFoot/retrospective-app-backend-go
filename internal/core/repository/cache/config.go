package repository_cache

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr        string        `envconfig:"ADDR"         required:"true"`
	Password    string        `envconfig:"PASSWORD"     required:"true"`
	User        string        `envconfig:"USER"         default:"default"`
	DB          int           `envconfig:"DB"           default:"0"`
	MaxRetries  int           `envconfig:"MAX_RETRIES"  default:"3"`
	DialTimeout time.Duration `envconfig:"DIAL_TIMEOUT" default:"5s"`
	Timeout     time.Duration `envconfig:"TIMEOUT"      default:"3s"`
}

func NewConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("REDIS", &config); err != nil {
		return Config{}, fmt.Errorf("process envonfig: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()

	if err != nil {
		err = fmt.Errorf("get redis connection pool config: %w", err)
		panic(err)
	}

	return config
}
