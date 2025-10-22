package postgres

import (
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/mmoldabe-dev/task-manager/config"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"  
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(cfg config.DatabaseConfig) error {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	m, err := migrate.New(
		"file://migration",
		dsn,
	)
	if err != nil {
		return fmt.Errorf("filed Migration file %w ", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}
	slog.Info("Database migrated successfully")
	return nil
}
