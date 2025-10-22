package domain

import "github.com/google/uuid"

type Tag struct {
	ID   uuid.UUID `db:"id" json:"id"`
	Name string    `db:"name_tags" json:"name_tags"`
}
