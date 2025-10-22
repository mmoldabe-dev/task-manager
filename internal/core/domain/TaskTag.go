package domain

import "github.com/google/uuid"

type TaskTag struct {
	TaskID uuid.UUID `db:"id_task" json:"id_task"`
	TagID  uuid.UUID `db:"id_tags" json:"id_tags"`
}
