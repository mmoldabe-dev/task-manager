package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/mmoldabe-dev/task-manager/config"
	"github.com/mmoldabe-dev/task-manager/internal/infrastructure/postgres"
	"github.com/mmoldabe-dev/task-manager/internal/infrastructure/rabbit"
	"github.com/mmoldabe-dev/task-manager/internal/infrastructure/redis"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Error load config")
		os.Exit(1)
	}
	fmt.Println(cfg)
	db, err := postgres.NewPostgresClient(ctx, cfg.Database)
	if err != nil {
		slog.Error("Error connecting to PostgreSQL.. ")
		os.Exit(1)
	}
	defer db.Close()

	redisdb, err := redis.RedisConnect(ctx, cfg.Redis)
	if err != nil {
		slog.Error("Error connecting to Redis..")
		os.Exit(1)
	}
	defer redisdb.Close()

	slog.Info("Connected to PostgreSQL successfully!")

	rabbitConn, rabbitCh, err := rabbit.RabbitClient(cfg.RabbitMQ)
	if err != nil {
		slog.Error("Error connecting to Rabbit..", "err", err)
		os.Exit(1)
	}

	defer rabbitCh.Close()
	defer rabbitConn.Close()

	if err := postgres.RunMigrations(cfg.Database); err != nil {
		slog.Error("Migration failed", "err", err)
		os.Exit(1)
	}

}
