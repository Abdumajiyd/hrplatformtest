package postgres

import (
	"time"

	"hrplatform/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CompanyRepository interface {
	CreateCompany(companyCreate models.CreateCompany) (models.Company, error)
	GetCompanyByID(id string) (models.Company, error)
	GetAllCompanies() ([]models.Company, error)
	UpdateCompany(companyUpdate models.UpdateCompany) error
	DeleteCompany(id string) error
}

type PostgresCompanyRepository struct {
	DB *sqlx.DB
}

func (r *PostgresCompanyRepository) CreateCompany(companyCreate models.CreateCompany) (models.Company, error) {
	company := models.Company{
		ID:        uuid.New(),
		Name:      companyCreate.Name,
		Location:  companyCreate.Location,
		Workers:   companyCreate.Workers,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: 0,
	}

	query := `INSERT INTO companies (id, name, location, workers, created_at, updated_at, deleted_at) 
	          VALUES (:id, :name, :location, :workers, :created_at, :updated_at, :deleted_at)`
	_, err := r.DB.NamedExec(query, company)
	if err != nil {
		return models.Company{}, err
	}

	var createdCompany models.Company
	err = r.DB.Get(&createdCompany, "SELECT * FROM companies WHERE id = $1", company.ID)
	if err != nil {
		return models.Company{}, err
	}

	return createdCompany, nil
}


func (r *PostgresCompanyRepository) GetCompanyByID(id string) (models.Company, error) {
	var company models.Company
	query := `SELECT * FROM companies WHERE id = $1`
	err := r.DB.Get(&company, query, id)
	if err != nil {
		return models.Company{}, err
	}
	return company, nil
}

func (r *PostgresCompanyRepository) GetAllCompanies() ([]models.Company, error) {
	var companies []models.Company
	query := `SELECT * FROM companies WHERE deleted_at = 0` // Exclude deleted companies
	err := r.DB.Select(&companies, query)
	if err != nil {
		return nil, err
	}
	return companies, nil
}
func (r *PostgresCompanyRepository) UpdateCompany(companyUpdate models.UpdateCompany) error {
	query := `
		UPDATE companies
		SET name = :name,
			location = :location,
			workers = :workers,
			updated_at = NOW()
		WHERE id = :id AND deleted_at = 0
	`

	_, err := r.DB.NamedExec(query, companyUpdate)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresCompanyRepository) DeleteCompany(id string) error {
	query := `DELETE FROM vacancies WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
