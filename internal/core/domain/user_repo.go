package domain

import (
	"context" // ⬅️ Обязателен для всех внешних операций

	"github.com/google/uuid"
)

// UserRepository - Порт для работы с данными пользователей.
type UserRepository interface {
	// В репозитории ТОЛЬКО операции с данными. Логика "Логин" здесь не нужна.

	CreateUser(ctx context.Context, user *User) error

	// Метод для аутентификации - нужен Use Case для сравнения хэша пароля
	GetUserByEmail(ctx context.Context, email string) (*User, error)

	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)

	GetAllUsers(ctx context.Context) ([]*User, error)

	UpdateUser(ctx context.Context, user *User) error

	// Мы не удаляем, а деактивируем пользователя
	DeactivateUser(ctx context.Context, id uuid.UUID) error
}
