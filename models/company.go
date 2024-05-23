package models

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Location  string    `db:"location" json:"location"`
	Workers   int       `db:"workers" json:"workers"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt int64     `db:"deleted_at" json:"deleted_at"`
}
type CreateCompany struct {
	Name      string    `db:"name" json:"name"`
	Location  string    `db:"location" json:"location"`
	Workers   int       `db:"workers" json:"workers"`
}
type UpdateCompany struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Location  string    `db:"location" json:"location"`
	Workers   int       `db:"workers" json:"workers"`
}
