package domain

import (
	"context"

	"github.com/google/uuid"
)

// TaskRepository - Порт для работы с задачами.
type TaskRepository interface {
	CreateTask(ctx context.Context, task *Task) error

	GetTaskByID(ctx context.Context, id uuid.UUID) (*Task, error)

	// Используем пагинацию/фильтрацию - сеньорский подход
	GetTasksByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*Task, error)

	UpdateTask(ctx context.Context, task *Task) error

	// Операции со связующей таблицей (Many-to-Many) живут здесь!
	AddTagToTask(ctx context.Context, taskID, tagID uuid.UUID) error
	RemoveTagFromTask(ctx context.Context, taskID, tagID uuid.UUID) error

	// Вместо DeletedTask используем мягкое удаление (is_deleted)
	SoftDeleteTask(ctx context.Context, id uuid.UUID) error
}
