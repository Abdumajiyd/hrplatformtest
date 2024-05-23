package postgres

import (
	"fmt"
	"hrplatform/models"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ResumeRepository interface {
	CreateResume(resumeCreate models.CreateResume) (models.Resume, error)
	GetResumeByID(id uuid.UUID) (models.Resume, error)
	GetResumesByUserID(userID uuid.UUID) ([]models.Resume, error)
	GetAllResumes(filter map[string]interface{}) ([]models.ResumeWithUser, error)
	UpdateResume(resumeUpdate models.UpdateResume) error
	DeleteResume(id string) error
}

type PostgresResumeRepository struct {
	DB *sqlx.DB
}

func (r *PostgresResumeRepository) CreateResume(resumeCreate models.CreateResume) (models.Resume, error) {
	resume := models.Resume{
		ID:          uuid.New(),
		Position:    resumeCreate.Position,
		Experience:  resumeCreate.Experience,
		Description: resumeCreate.Description,
		UserID:      resumeCreate.UserID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   0,
	}

	query := `INSERT INTO resumes (id, position, experience, description, user_id, created_at, updated_at, deleted_at) 
              VALUES (:id, :position, :experience, :description, :user_id, :created_at, :updated_at, :deleted_at)`

	_, err := r.DB.NamedExec(query, resume)
	if err != nil {
		return models.Resume{}, err
	}

	return resume, nil
}

func (r *PostgresResumeRepository) GetResumeByID(id uuid.UUID) (models.Resume, error) {
	var resume models.Resume
	query := `SELECT * FROM resumes WHERE id = $1`
	err := r.DB.Get(&resume, query, id)
	if err != nil {
		return models.Resume{}, err
	}
	return resume, nil
}

func (r *PostgresResumeRepository) GetResumesByUserID(userID uuid.UUID) ([]models.Resume, error) {
	var resumes []models.Resume
	query := `SELECT * FROM resumes WHERE user_id = $1`
	err := r.DB.Select(&resumes, query, userID)
	if err != nil {
		return nil, err
	}
	return resumes, nil
}
func (r *PostgresResumeRepository) GetAllResumes(filter map[string]interface{}) ([]models.ResumeWithUser, error) {
	var resumes []models.ResumeWithUser


	query := `
        SELECT r.*, u.name AS user_name, u.email AS user_email 
        FROM resumes r
        JOIN users u ON r.user_id = u.id
        WHERE r.deleted_at = 0
    `

	var conditions []string
	var args []interface{}

	for key, value := range filter {
		switch key {
		case "position":
			conditions = append(conditions, fmt.Sprintf("r.position ILIKE '%%%s%%'", value))
		case "min_exp":
			conditions = append(conditions, fmt.Sprintf("r.experience >= %d", value))
		}
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}


	err := r.DB.Select(&resumes, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get all resumes: %w", err)
	}

	return resumes, nil
}

// func (r *PostgresResumeRepository) GetAllResumes(filter map[string]interface{}) ([]models.ResumeWithUser, error) {
// 	var resumes []models.ResumeWithUser

// 	// Base query with join
// 	query := `
//         SELECT r.*, u.name AS user_name, u.email AS user_email
//         FROM resumes r
//         JOIN users u ON r.user_id = u.id
//         WHERE r.deleted_at = 0
//     `

// 	var conditions []string
// 	var args []interface{}

// 	// Adding conditions based on filters
// 	for key, value := range filter {
// 		switch key {
// 		case "position":
// 			conditions = append(conditions, "r.position ILIKE '%' || ? || '%'")
// 			args = append(args, value.(string))
// 		case "min_exp":
// 			conditions = append(conditions, "r.experience >= ?")
// 			args = append(args, value.(int))
// 		}
// 	}

// 	// Append conditions to the query if any
// 	if len(conditions) > 0 {
// 		query += " AND " + strings.Join(conditions, " AND ")
// 	}

// 	// Execute the query
// 	err := r.DB.Select(&resumes, query, args...)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get all resumes: %w", err)
// 	}

//		return resumes, nil
//	}
func (r *PostgresResumeRepository) UpdateResume(resumeUpdate models.UpdateResume) error {
	fields := make(map[string]interface{})

	if resumeUpdate.Position != "" {
		fields["position"] = resumeUpdate.Position
	}
	if resumeUpdate.Experience != 0 {
		fields["experience"] = resumeUpdate.Experience
	}
	if resumeUpdate.Description != "" {
		fields["description"] = resumeUpdate.Description
	}
	if resumeUpdate.UserID != uuid.Nil {
		fields["user_id"] = resumeUpdate.UserID
	}

	if len(fields) == 0 {
		return nil
	}

	setClauses := []string{}
	for key := range fields {
		setClauses = append(setClauses, fmt.Sprintf("%s = :%s", key, key))
	}
	setQuery := strings.Join(setClauses, ", ")

	query := fmt.Sprintf("UPDATE resumes SET %s, updated_at = NOW() WHERE id = :id", setQuery)
	fields["id"] = uuid.MustParse(resumeUpdate.ID)

	_, err := r.DB.NamedExec(query, fields)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresResumeRepository) DeleteResume(id string) error {
	query := `DELETE FROM resumes WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
