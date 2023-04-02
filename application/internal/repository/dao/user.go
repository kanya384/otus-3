package dao

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID `db:"id"`
	UserName   string    `db:"user_name"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	Email      string    `db:"email"`
	Phone      string    `db:"phone"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`
}

var UserColumns = []string{
	"id",
	"user_name",
	"first_name",
	"last_name",
	"email",
	"phone",
	"created_at",
	"modified_at",
}
