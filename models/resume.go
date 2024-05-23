package models

import (
	"time"

	"github.com/google/uuid"
)

type Resume struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Position    string    `db:"position" json:"position"`
	Experience  int       `db:"experience" json:"experience"`
	Description string    `db:"description" json:"description"`
	UserID      uuid.UUID `db:"user_id" json:"user_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt   int64     `db:"deleted_at" json:"deleted_at"`
}

type CreateResume struct {
	Position    string    `db:"position" json:"position"`
	Experience  int       `db:"experience" json:"experience"`
	Description string    `db:"description" json:"description"`
	UserID      uuid.UUID `db:"user_id" json:"user_id"`
}

type UpdateResume struct {
	ID          string    `db:"id" json:"id"`
	Position    string    `db:"position" json:"position"`
	Experience  int       `db:"experience" json:"experience"`
	Description string    `db:"description" json:"description"`
	UserID      uuid.UUID `db:"user_id" json:"user_id"`
}

type ResumeWithUser struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Position    string    `db:"position" json:"position"`
	Experience  int       `db:"experience" json:"experience"`
	Description string    `db:"description" json:"description"`
	UserID      uuid.UUID `db:"user_id" json:"user_id"`
	UserName    string    `db:"user_name" json:"user_name"`
	UserEmail   string    `db:"user_email" json:"user_email"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt   int64     `db:"deleted_at" json:"deleted_at"`
}
