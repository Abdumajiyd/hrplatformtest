package models

import (
	"time"

	"github.com/google/uuid"
)

type Vacancy struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Position    string    `db:"position" json:"position"`
	MinExp      int       `db:"min_exp" json:"min_exp"`
	CompanyID   uuid.UUID `db:"company_id" json:"company_id"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt   int64     `db:"deleted_at" json:"deleted_at"`
}
type CreateVacancy struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Position    string    `db:"position" json:"position"`
	MinExp      int       `db:"min_exp" json:"min_exp"`
	CompanyID   uuid.UUID `db:"company_id" json:"company_id"`
	Description string    `db:"description" json:"description"`
}

type UpdateVacancy struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Position    string    `db:"position" json:"position"`
	MinExp      int       `db:"min_exp" json:"min_exp"`
	CompanyID   uuid.UUID `db:"company_id" json:"company_id"`
}