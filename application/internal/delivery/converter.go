package delivery

import (
	jsonUser "otus/internal/delivery/user"
	"otus/internal/domain"
)

func (d *Delivery) toUserResponse(user *domain.User) *jsonUser.UserResponse {
	return &jsonUser.UserResponse{
		Id:         user.Id(),
		UserName:   user.UserName(),
		FirstName:  user.FirstName(),
		LastName:   user.LastName(),
		Email:      user.Email(),
		Phone:      user.Phone(),
		CreatedAt:  user.CreatedAt(),
		ModifiedAt: user.ModifiedAt(),
	}
}
