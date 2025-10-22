package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id           uuid.UUID `db:"id" json:"id"`
	UserId      uuid.UUID `db:"user_id" json:"user_id"`
	Title        string    `db:"title" json:"title"`
	Description   string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	CompletedAt *time.Time `db:"completed_at" json:"completed_at"`
	Priority     int       `db:"priority" json:"priority"`
	IsDeleted   bool      `db:"is_deleted" json:"is_deleted"`
}
	