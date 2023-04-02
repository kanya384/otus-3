package service

import (
	"context"
	"otus/internal/domain"
	"time"

	"github.com/google/uuid"
)

func (s *service) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	user, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		s.logger.Error("CreateUser - error creating user: %s", err.Error())
	}
	return user, err
}

func (s *service) UpdateUser(ctx context.Context, updateUserParams *UserUpdateParams) (*domain.User, error) {
	upFn := func(oldUser *domain.User) (*domain.User, error) {
		return domain.NewUserWithId(oldUser.Id(), updateUserParams.UserName, updateUserParams.FirstName, updateUserParams.LastName, updateUserParams.Email, updateUserParams.Phone, oldUser.CreatedAt(), time.Now())
	}
	user, err := s.repository.UpdateUser(ctx, updateUserParams.Id, upFn)
	if err != nil {
		s.logger.Error("UpdateUser - error updating user: %s", err.Error())
	}
	return user, err
}

func (s *service) DeleteUser(ctx context.Context, id uuid.UUID) (err error) {
	err = s.repository.DeleteUser(ctx, id)
	if err != nil {
		s.logger.Error("DeleteUser - error deleting user: %s", err.Error())
	}
	return
}

func (s *service) ReadUserById(ctx context.Context, id uuid.UUID) (user *domain.User, err error) {
	user, err = s.repository.ReadUserById(ctx, id)
	if err != nil {
		s.logger.Error("ReadUserById - error reading user by id: %s", err.Error())
	}
	return
}
