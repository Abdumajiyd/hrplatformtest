package postgres

import (
	"fmt"
	"hrplatform/models"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type InterviewRepository interface {
	CreateInterview(interviewCreate models.CreateInterview) (models.Interview, error)
	GetInterviewByID(id uuid.UUID) (models.Interview, error)
	GetInterviewsByUserID(userID uuid.UUID) ([]models.Interview, error)
	GetAllInterviews(filter map[string]interface{}) ([]models.Interview, error)
	UpdateInterview(interviewUpdate models.UpdateInterview) error
	DeleteInterview(id uuid.UUID) error
}

type PostgresInterviewRepository struct {
	DB *sqlx.DB
}

func (r *PostgresInterviewRepository) CreateInterview(interviewCreate models.CreateInterview) (models.Interview, error) {
	
    var userAge int
    err := r.DB.Get(&userAge, `SELECT EXTRACT(YEAR FROM AGE(birthday)) FROM users WHERE id = $1`, interviewCreate.UserID)
    if err != nil {
        return models.Interview{}, err
    }
    if userAge < 18 {
        return models.Interview{}, fmt.Errorf("user must be at least 18 years old")
    }

    var resumePosition, vacancyPosition string
    err = r.DB.Get(&resumePosition, `SELECT position FROM resumes WHERE user_id = $1`, interviewCreate.UserID)
    if err != nil {
        return models.Interview{}, err
    }
    err = r.DB.Get(&vacancyPosition, `SELECT position FROM vacancies WHERE id = $1`, interviewCreate.VacancyID)
    if err != nil {
        return models.Interview{}, err
    }
    if resumePosition != vacancyPosition {
        return models.Interview{}, fmt.Errorf("position in resume and vacancy must match")
    }
	interviewDate, err := time.Parse("2006-01-02 15:04:05", interviewCreate.InterviewDate)
	if err != nil {
		return models.Interview{}, err
	}

	interview := models.Interview{
		ID:            uuid.New(),
		UserID:        interviewCreate.UserID,
		VacancyID:     interviewCreate.VacancyID,
		RecruiterID:   interviewCreate.RecruiterID,
		InterviewDate: interviewDate,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     0,
	}

	query := `INSERT INTO interviews (id, user_id, vacancy_id, recruiter_id, interview_date, created_at, updated_at, deleted_at) 
              VALUES (:id, :user_id, :vacancy_id, :recruiter_id, :interview_date, :created_at, :updated_at, :deleted_at)`

	_, err = r.DB.NamedExec(query, interview)
	if err != nil {
		return models.Interview{}, err
	}

	return interview, nil
}

func (r *PostgresInterviewRepository) GetInterviewByID(id uuid.UUID) (models.Interview, error) {
	var interview models.Interview
	query := `SELECT * FROM interviews WHERE id = $1`
	err := r.DB.Get(&interview, query, id)
	if err != nil {
		return models.Interview{}, err
	}
	return interview, nil
}

func (r *PostgresInterviewRepository) GetInterviewsByUserID(userID uuid.UUID) ([]models.Interview, error) {
	var interviews []models.Interview
	query := `SELECT * FROM interviews WHERE user_id = $1`
	err := r.DB.Select(&interviews, query, userID)
	if err != nil {
		return nil, err
	}
	return interviews, nil
}

func (r *PostgresInterviewRepository) GetAllInterviews(filter map[string]interface{}) ([]models.Interview, error) {
	var interviews []models.Interview
	baseQuery := `SELECT * FROM interviews WHERE deleted_at = 0`

	var conditions []string
	var args []interface{}

	for key, value := range filter {
		switch key {
		case "company_id":
			conditions = append(conditions, "recruiter_id IN (SELECT id FROM recruiters WHERE company_id = $1)")
			args = append(args, value)
		case "position":
			conditions = append(conditions, "vacancy_id IN (SELECT id FROM vacancies WHERE position ILIKE ?)")
			args = append(args, "%"+value.(string)+"%")
		case "experience":
			conditions = append(conditions, "user_id IN (SELECT user_id FROM resumes WHERE experience >= ?)")
			args = append(args, value)
		}
	}

	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	err := r.DB.Select(&interviews, baseQuery, args...)
	if err != nil {
		return nil, err
	}

	return interviews, nil
}

func (r *PostgresInterviewRepository) UpdateInterview(interviewUpdate models.UpdateInterview) error {
	interviewDate, err := time.Parse("2006-01-02 15:04:05", interviewUpdate.InterviewDate)
	if err != nil {
		return err
	}

	fields := map[string]interface{}{
		"user_id":        interviewUpdate.UserID,
		"vacancy_id":     interviewUpdate.VacancyID,
		"recruiter_id":   interviewUpdate.RecruiterID,
		"interview_date": interviewDate,
	}

	setClauses := []string{}
	for key := range fields {
		setClauses = append(setClauses, fmt.Sprintf("%s = :%s", key, key))
	}
	setQuery := strings.Join(setClauses, ", ")

	query := fmt.Sprintf("UPDATE interviews SET %s, updated_at = NOW() WHERE id = :id", setQuery)
	fields["id"] = interviewUpdate.ID

	_, err = r.DB.NamedExec(query, fields)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresInterviewRepository) DeleteInterview(id uuid.UUID) error {
	query := `DELETE FROM interviews WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
