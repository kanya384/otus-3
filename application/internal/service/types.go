package service

import "github.com/google/uuid"

type UserUpdateParams struct {
	Id        uuid.UUID
	UserName  string
	FirstName string
	LastName  string
	Email     string
	Phone     string
}
