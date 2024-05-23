package postgres
////////////////
import (
	"fmt"
	"hrplatform/models"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type VacancyRepository interface {
	CreateVacancy(vacancyCreate models.CreateVacancy) (models.Vacancy, error)
	GetVacancyByID(id uuid.UUID) (models.Vacancy, error)
	GetAllVacancies(filter map[string]interface{}) ([]models.Vacancy, error)
	UpdateVacancy(vacancyUpdate models.UpdateVacancy) error
	DeleteVacancy(id uuid.UUID) error
	CheckCompanyExists(companyID uuid.UUID) (bool, error)
}

type PostgresVacancyRepository struct {
	DB *sqlx.DB
}

func (r *PostgresVacancyRepository) CreateVacancy(vacancyCreate models.CreateVacancy) (models.Vacancy, error) {
	vacancy := models.Vacancy{
		ID:          uuid.New(),
		Name:        vacancyCreate.Name,
		Position:    vacancyCreate.Position,
		MinExp:      vacancyCreate.MinExp,
		CompanyID:   vacancyCreate.CompanyID,
		Description: vacancyCreate.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   0,
	}

	query := `INSERT INTO vacancies (id, name, position, min_exp, company_id, description, created_at, updated_at, deleted_at) 
              VALUES (:id, :name, :position, :min_exp, :company_id, :description, :created_at, :updated_at, :deleted_at)`

	_, err := r.DB.NamedExec(query, vacancy)
	if err != nil {
		return models.Vacancy{}, err
	}

	return vacancy, nil
}

func (r *PostgresVacancyRepository) GetVacancyByID(id uuid.UUID) (models.Vacancy, error) {
	var vacancy models.Vacancy
	query := `SELECT * FROM vacancies WHERE id = $1`
	err := r.DB.Get(&vacancy, query, id)
	if err != nil {
		return models.Vacancy{}, err
	}
	return vacancy, nil
}

func (r *PostgresVacancyRepository) GetAllVacancies(filter map[string]interface{}) ([]models.Vacancy, error) {
	var vacancies []models.Vacancy
	var conditions []string
	var params []interface{}

	for key, value := range filter {
		switch key {
		case "position":
			conditions = append(conditions, "position ILIKE ?")
			params = append(params, "%"+value.(string)+"%")
		case "min_exp":
			conditions = append(conditions, "min_exp >= ?")
			params = append(params, value.(int))
		case "company_id":
			conditions = append(conditions, "company_id = ?")
			params = append(params, value.(uuid.UUID))
		}
	}

	query := "SELECT * FROM vacancies"
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	err := r.DB.Select(&vacancies, query, params...)
	if err != nil {
		return nil, err
	}
	return vacancies, nil
}

func (r *PostgresVacancyRepository) UpdateVacancy(vacancyUpdate models.UpdateVacancy) error {
	fields := make(map[string]interface{})

	if vacancyUpdate.Name != "" {
		fields["name"] = vacancyUpdate.Name
	}
	if vacancyUpdate.Position != "" {
		fields["position"] = vacancyUpdate.Position
	}
	if vacancyUpdate.MinExp != 0 {
		fields["min_exp"] = vacancyUpdate.MinExp
	}
	if vacancyUpdate.CompanyID != uuid.Nil {
		fields["company_id"] = vacancyUpdate.CompanyID
	}

	if len(fields) == 0 {
		return nil
	}

	setClauses := []string{}
	for key := range fields {
		setClauses = append(setClauses, fmt.Sprintf("%s = :%s", key, key))
	}
	setQuery := strings.Join(setClauses, ", ")

	query := fmt.Sprintf("UPDATE vacancies SET %s, updated_at = NOW() WHERE id = :id", setQuery)
	fields["id"] = vacancyUpdate.ID

	_, err := r.DB.NamedExec(query, fields)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresVacancyRepository) DeleteVacancy(id uuid.UUID) error {
	query := `DELETE FROM vacancies WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresVacancyRepository) CheckCompanyExists(companyID uuid.UUID) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM companies WHERE LOWER(id::text) = $1)`
	err := r.DB.Get(&exists, query, strings.ToLower(companyID.String()))
	return exists, err
}
