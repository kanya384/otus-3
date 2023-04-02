package repository

import (
	"otus/internal/domain"
	"otus/internal/repository/dao"
)

func (r *Repository) toDomainUser(user *dao.User) (*domain.User, error) {
	return domain.NewUserWithId(user.Id, user.UserName, user.FirstName, user.LastName, user.Email, user.Phone, user.CreatedAt, user.ModifiedAt)
}
