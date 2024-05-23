package models

import (
	"time"

	"github.com/google/uuid"
)

type Interview struct {
	ID            uuid.UUID `db:"id" json:"id"`
	UserID        uuid.UUID `db:"user_id" json:"user_id"`
	VacancyID     uuid.UUID `db:"vacancy_id" json:"vacancy_id"`
	RecruiterID   uuid.UUID `db:"recruiter_id" json:"recruiter_id"`
	InterviewDate time.Time `db:"interview_date" json:"interview_date"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt     int64     `db:"deleted_at" json:"deleted_at"`
}

type CreateInterview struct {
	UserID        uuid.UUID `db:"user_id" json:"user_id"`
	VacancyID     uuid.UUID `db:"vacancy_id" json:"vacancy_id"`
	RecruiterID   uuid.UUID `db:"recruiter_id" json:"recruiter_id"`
	InterviewDate string    `db:"interview_date" json:"interview_date"`
}

type UpdateInterview struct {
	ID            uuid.UUID `db:"id" json:"id"`
	UserID        uuid.UUID `db:"user_id" json:"user_id"`
	VacancyID     uuid.UUID `db:"vacancy_id" json:"vacancy_id"`
	RecruiterID   uuid.UUID `db:"recruiter_id" json:"recruiter_id"`
	InterviewDate string    `db:"interview_date" json:"interview_date"`
}
