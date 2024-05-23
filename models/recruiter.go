package models

import (
	"time"

	"github.com/google/uuid"
)

type Recruiter struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Email       string    `db:"email" json:"email"`
	PhoneNumber string    `db:"phone_number" json:"phone_number"`
	Birthday    time.Time `db:"birthday" json:"birthday"`
	Gender      string    `db:"gender" json:"gender"`
	CompanyID   uuid.UUID `db:"company_id" json:"company_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt   int64     `db:"deleted_at" json:"deleted_at"`
}

type GetRecruiter struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Email       string    `db:"email" json:"email"`
	PhoneNumber string    `db:"phone_number" json:"phone_number"`
	Birthday    time.Time `db:"birthday" json:"birthday"`
	Gender      string    `db:"gender" json:"gender"`
	CompanyID   uuid.UUID `db:"company_id" json:"company_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt   int64     `db:"deleted_at" json:"deleted_at"`
}

type CreateRecruiter struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Birthday    string    `json:"birthday"` // Change to string to match the JSON input
	Gender      string    `json:"gender"`
	CompanyID   uuid.UUID `json:"company_id"`
}

type UpdateRecruiter struct {
	ID          uuid.UUID  `db:"id" json:"id"`
	Name        *string    `db:"name" json:"name"`
	Email       *string    `db:"email" json:"email"`
	PhoneNumber *string    `db:"phone_number" json:"phone_number"`
	Birthday    *string    `json:"birthday"` 
	Gender      *string    `json:"gender"`
	CompanyID   *uuid.UUID `json:"company_id"`
}
