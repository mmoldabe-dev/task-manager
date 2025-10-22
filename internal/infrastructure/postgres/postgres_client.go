package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
	"github.com/mmoldabe-dev/task-manager/config"
)

func NewPostgresClient(ctx context.Context, cfg config.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open PostgreSQL connection: %w", err)

	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)
	const maxAttempts = 5
	for i := 1; i <= maxAttempts; i++ {
		if err := db.PingContext(ctx); err == nil {
			slog.Info("Connected to PostgreSQL successfully")
			return db, nil
		}

		slog.Warn("Failed to connect to PostgreSQL, retrying...",
			slog.Int("attempt", i),
			slog.Int("max_attempts", maxAttempts),
		)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("could not connect to PostgreSQL after %d attempts", maxAttempts)

}
