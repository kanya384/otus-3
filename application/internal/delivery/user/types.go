package user

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	UserName  string `json:"userName" example:"johndoe589"`
	FirstName string `json:"firstName" example:"John"`
	LastName  string `json:"lastName" example:"Doe"`
	Email     string `json:"email" example:"bestjohn@doe.com"`
	Phone     string `json:"phone" example:"+71002003040"`
}

type UpdateUserRequest struct {
	Id        uuid.UUID `json:"id" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
	UserName  string    `json:"userName" example:"johndoe589"`
	FirstName string    `json:"firstName" example:"John"`
	LastName  string    `json:"lastName" example:"Doe"`
	Email     string    `json:"email" example:"bestjohn@doe.com"`
	Phone     string    `json:"phone" example:"+71002003040"`
}

type UserResponse struct {
	Id         uuid.UUID `json:"id" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
	UserName   string    `json:"userName" example:"johndoe589"`
	FirstName  string    `json:"firstName" example:"John"`
	LastName   string    `json:"lastName" example:"Doe"`
	Email      string    `json:"email" example:"bestjohn@doe.com"`
	Phone      string    `json:"phone" example:"+71002003040"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}
