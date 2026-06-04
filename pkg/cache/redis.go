package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Config holds the settings needed to connect to Redis.
type Config struct {
	Addr     string // host:port, e.g. "localhost:6379"
	Password string // empty if no auth
	DB       int    // logical database number (default 0)
}

// New creates a Redis client and verifies connectivity with a Ping.
// The caller is responsible for closing the returned client.
func New(ctx context.Context, cfg Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := client.Ping(pingCtx).Err(); err != nil {
		_ = client.Close()
		return nil, fmt.Errorf("ping redis: %w", err)
	}

	return client, nil
}
