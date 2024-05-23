package postgres

import (
	"fmt"
	"hrplatform/models"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RecruiterRepository interface {
	CreateRecruiter(recruiterCreate models.CreateRecruiter) (models.Recruiter, error)
	GetRecruiterByID(id string) (models.Recruiter, error)
	GetAllRecruiters(age int, gender string, companyID string) ([]models.Recruiter, error)
	UpdateRecruiter(recruiterUpdate models.UpdateRecruiter) error
	DeleteRecruiter(id string) error
	CheckCompanyExists(companyID uuid.UUID) (bool, error)
}

type PostgresRecruiterRepository struct {
	DB *sqlx.DB
}

func (r *PostgresRecruiterRepository) CreateRecruiter(recruiterCreate models.CreateRecruiter) (models.Recruiter, error) {
	// Check if company exists
	exists, err := r.CheckCompanyExists(recruiterCreate.CompanyID)
	if err != nil {
		return models.Recruiter{}, err
	}
	if !exists {
		return models.Recruiter{}, fmt.Errorf("company with ID %s does not exist", recruiterCreate.CompanyID)
	}

	// Parse birthday from string to time.Time
	birthday, err := time.Parse(time.RFC3339, recruiterCreate.Birthday)
	if err != nil {
		return models.Recruiter{}, fmt.Errorf("invalid birthday format")
	}

	recruiter := models.Recruiter{
		ID:          uuid.New(),
		Name:        recruiterCreate.Name,
		Email:       recruiterCreate.Email,
		PhoneNumber: recruiterCreate.PhoneNumber,
		Birthday:    birthday,
		Gender:      recruiterCreate.Gender,
		CompanyID:   recruiterCreate.CompanyID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   0,
	}

	query := `INSERT INTO recruiters (id, name, email, phone_number, birthday, gender, company_id, created_at, updated_at, deleted_at)
              VALUES (:id, :name, :email, :phone_number, :birthday, :gender, :company_id, :created_at, :updated_at, :deleted_at)`

	_, err = r.DB.NamedExec(query, recruiter)
	if err != nil {
		log.Printf("Error creating recruiter: %v\nQuery: %s", err, query)
		return models.Recruiter{}, err
	}

	return recruiter, nil
}

func (r *PostgresRecruiterRepository) CheckCompanyExists(companyID uuid.UUID) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM companies WHERE id = $1)`
	err := r.DB.Get(&exists, query, companyID)
	return exists, err
}

func (r *PostgresRecruiterRepository) GetRecruiterByID(id string) (models.Recruiter, error) {
	var recruiter models.Recruiter
	query := `SELECT * FROM recruiters WHERE id = $1`
	err := r.DB.Get(&recruiter, query, id)
	if err != nil {
		return models.Recruiter{}, err
	}
	return recruiter, nil
}

func (r *PostgresRecruiterRepository) GetAllRecruiters(age int, gender string, companyID string) ([]models.Recruiter, error) {
	var recruiters []models.Recruiter
	var filters []string
	var args []interface{}
	argCount := 1

	if age > 0 {
		filters = append(filters, fmt.Sprintf("EXTRACT(YEAR FROM AGE(birthday)) = $%d", argCount))
		args = append(args, age)
		argCount++
	}
	if gender != "" {
		filters = append(filters, fmt.Sprintf("gender = $%d", argCount))
		args = append(args, gender)
		argCount++
	}
	if companyID != "" {
		companyID = strings.TrimSpace(companyID)
		filters = append(filters, fmt.Sprintf("company_id = $%d", argCount))
		args = append(args, companyID)
		argCount++
	}

	filterQuery := strings.Join(filters, " AND ")
	if filterQuery != "" {
		filterQuery = "WHERE " + filterQuery
	}

	query := fmt.Sprintf("SELECT * FROM recruiters %s", filterQuery)
	err := r.DB.Select(&recruiters, query, args...)
	if err != nil {
		return nil, err
	}
	return recruiters, nil
}


func (r *PostgresRecruiterRepository) UpdateRecruiter(recruiter models.UpdateRecruiter) error {
    fields := []string{}
    params := map[string]interface{}{"id": recruiter.ID}

    if recruiter.Name != nil {
        fields = append(fields, "name = :name")
        params["name"] = *recruiter.Name
    }
    if recruiter.Email != nil {
        fields = append(fields, "email = :email")
        params["email"] = *recruiter.Email
    }
    if recruiter.PhoneNumber != nil {
        fields = append(fields, "phone_number = :phone_number")
        params["phone_number"] = *recruiter.PhoneNumber
    }
    if recruiter.Birthday != nil {
        birthday, err := time.Parse(time.RFC3339, *recruiter.Birthday)
        if err != nil {
            return fmt.Errorf("invalid date format for birthday")
        }
        params["birthday"] = birthday
        fields = append(fields, "birthday = :birthday")
    }
    if recruiter.Gender != nil {
        fields = append(fields, "gender = :gender")
        params["gender"] = *recruiter.Gender
    }
    if recruiter.CompanyID != nil {
        fields = append(fields, "company_id = :company_id")
        params["company_id"] = *recruiter.CompanyID
    }

    if len(fields) == 0 {
        return fmt.Errorf("no fields to update")
    }

    fields = append(fields, "updated_at = NOW()")
    query := fmt.Sprintf("UPDATE recruiters SET %s WHERE id = :id", strings.Join(fields, ", "))

    _, err := r.DB.NamedExec(query, params)
    return err
}


func (r *PostgresRecruiterRepository) DeleteRecruiter(id string) error {
	query := `DELETE FROM recruiters WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
