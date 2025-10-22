package redis

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/mmoldabe-dev/task-manager/config"
	"github.com/redis/go-redis/v9"
)

func RedisConnect(ctx context.Context, cfg config.RedisConfig) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	slog.Info("Syccesfully redis connect!")
	return client, nil
}
