package domain

import (
	"context"

	"github.com/google/uuid"
)

// TagRepository - Порт для работы с метками.
type TagRepository interface {
	CreateTag(ctx context.Context, tag *Tag) error

	GetTagByID(ctx context.Context, id uuid.UUID) (*Tag, error)

	// Получение меток по их названию (для проверки уникальности)
	GetTagByName(ctx context.Context, name string) (*Tag, error)

	GetAllTags(ctx context.Context) ([]*Tag, error)
}
