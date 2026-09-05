package repository_cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Get(ctx context.Context, key string, dest any) error
	Delete(ctx context.Context, key string) error
	Close() error
}

type ConnectionCache struct {
	*redis.Client
}

func NewConnectionCache(ctx context.Context, config Config) (*ConnectionCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         config.Addr,
		Password:     config.Password,
		DB:           config.DB,
		Username:     config.User,
		MaxRetries:   config.MaxRetries,
		DialTimeout:  config.DialTimeout,
		ReadTimeout:  config.Timeout,
		WriteTimeout: config.Timeout,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		fmt.Printf("failed to connect to redis server: %s\n", err.Error())
		return nil, err
	}

	return &ConnectionCache{Client: client}, nil
}

func (c *ConnectionCache) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("unable to set data to cache: %w", err)
	}
	return c.Client.Set(ctx, key, data, ttl).Err()
}

func (c *ConnectionCache) Get(ctx context.Context, key string, dest any) error {
	data, err := c.Client.Get(ctx, key).Bytes()
	if err != nil {
		return fmt.Errorf("unable to get data from cache: %w", err)
	}

	return json.Unmarshal(data, dest)
}

func (c *ConnectionCache) Delete(ctx context.Context, key string) error {
	if err := c.Client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("unable to delete data from cache: %w", err)
	}

	return nil
}
