package service

import (
	"context"
	"otus/internal/domain"

	"github.com/google/uuid"
)

type Service interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	UpdateUser(ctx context.Context, updateUserParams *UserUpdateParams) (*domain.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) (err error)
	ReadUserById(ctx context.Context, id uuid.UUID) (user *domain.User, err error)
}
