package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	id         uuid.UUID
	userName   string
	firstName  string
	lastName   string
	email      string
	phone      string
	createdAt  time.Time
	modifiedAt time.Time
}

func (u User) Id() uuid.UUID {
	return u.id
}

func (u User) UserName() string {
	return u.userName
}

func (u User) FirstName() string {
	return u.firstName
}

func (u User) LastName() string {
	return u.lastName
}

func (u User) Email() string {
	return u.email
}

func (u User) Phone() string {
	return u.phone
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

func (u User) ModifiedAt() time.Time {
	return u.modifiedAt
}

func NewUserWithId(
	id uuid.UUID,
	userName string,
	firstName string,
	lastName string,
	email string,
	phone string,
	createdAt time.Time,
	modifiedAt time.Time,
) (*User, error) {
	return &User{
		id:         id,
		userName:   userName,
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		phone:      phone,
		createdAt:  createdAt,
		modifiedAt: modifiedAt,
	}, nil
}

func NewUser(
	userName string,
	firstName string,
	lastName string,
	email string,
	phone string,
) (*User, error) {
	return &User{
		id:         uuid.New(),
		userName:   userName,
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		phone:      phone,
		createdAt:  time.Now(),
		modifiedAt: time.Now(),
	}, nil
}
