package postgres

import (
	"database/sql"

	"github.com/mmoldabe-dev/task-manager/internal/core/domain"
)

type PostgresTaskRepository struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{db: db}
}

var _ domain.TaskRepository = (*PostgresTaskRepository)(nil)

func(r * PostgresTaskRepository)CreateTask
