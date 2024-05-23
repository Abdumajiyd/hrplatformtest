package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Email       string    `db:"email" json:"email"`
	PhoneNumber string    `db:"phone_number" json:"phone_number"`
	Birthday    time.Time `db:"birthday" json:"birthday"`
	Gender      string    `db:"gender" json:"gender"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt   int64     `db:"deleted_at" json:"deleted_at"`
}

type UserCreate struct {
	Name        string `db:"name" json:"name"`
	Email       string `db:"email" json:"email"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	Birthday    string `db:"birthday" json:"birthday"`
	Gender      string `db:"gender" json:"gender"`
}

type UserUpdate struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Email       string    `db:"email" json:"email"`
	PhoneNumber string    `db:"phone_number" json:"phone_number"`
	Birthday    string    `db:"birthday" json:"birthday"`
	Gender      string    `db:"gender" json:"gender"`
}
