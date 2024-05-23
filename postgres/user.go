package postgres

////////////////
import (
	"fmt"
	"hrplatform/models"
	"strings"
	"time"
	// "log/slog"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(userCreate models.UserCreate) (models.User, error)
	GetUserByID(id string) (models.User, error)
	GetAllUsers(filters map[string]interface{}) ([]models.User, error)
	UpdateUser(userUpdate models.UserUpdate) error
	DeleteUser(id string) error
	GetUserInterviews(userID uuid.UUID) ([]models.Interview, error)
	GetUserResume(userID uuid.UUID) ([]models.Resume, error)
}

type PostgresUserRepository struct {
	DB *sqlx.DB
}

func (r *PostgresUserRepository) CreateUser(userCreate models.UserCreate) (models.User, error) {
	birthday, err := time.Parse("2006-01-02", userCreate.Birthday)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		ID:          uuid.New(),
		Name:        userCreate.Name,
		Email:       userCreate.Email,
		PhoneNumber: userCreate.PhoneNumber,
		Birthday:    birthday,
		Gender:      userCreate.Gender,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   0,
	}

	query := `INSERT INTO users (id, name, email, phone_number, birthday, gender, created_at, updated_at, deleted_at) 
              VALUES (:id, :name, :email, :phone_number, :birthday, :gender, :created_at, :updated_at, :deleted_at)`

	_, err = r.DB.NamedExec(query, user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *PostgresUserRepository) GetUserByID(id string) (models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE id = $1`
	err := r.DB.Get(&user, query, id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *PostgresUserRepository) GetAllUsers(filters map[string]interface{}) ([]models.User, error) {
	var users []models.User
	query := `SELECT * FROM users WHERE deleted_at = 0`
	params := []interface{}{}
	counter := 1
	// // Filtrlar uchun shartlarni queryga qo'shish
	if age, ok := filters["age"].(int); ok {
		query += fmt.Sprintf(" AND EXTRACT(YEAR FROM AGE(birthday)) = $%d", counter)
		params = append(params, age)
		counter++
	}

	if gender, ok := filters["gender"].(string); ok {
		query += fmt.Sprintf(" AND gender = $%d", counter)
		params = append(params, gender)
		counter++
	}

	err := r.DB.Select(&users, query, params...)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *PostgresUserRepository) UpdateUser(userUpdate models.UserUpdate) error {
	fields := make(map[string]interface{})

	if userUpdate.Name != "" {
		fields["name"] = userUpdate.Name
	}
	if userUpdate.Email != "" {
		fields["email"] = userUpdate.Email
	}
	if userUpdate.PhoneNumber != "" {
		fields["phone_number"] = userUpdate.PhoneNumber
	}
	if userUpdate.Birthday != "" {
		birthday, err := time.Parse("2006-01-02", userUpdate.Birthday)
		if err != nil {
			return err
		}
		fields["birthday"] = birthday
	}
	if userUpdate.Gender != "" {
		fields["gender"] = userUpdate.Gender
	}

	if len(fields) == 0 {
		return nil
	}

	setClauses := []string{}
	for key := range fields {
		setClauses = append(setClauses, fmt.Sprintf("%s = :%s", key, key))
	}
	setQuery := strings.Join(setClauses, ", ")

	query := fmt.Sprintf("UPDATE users SET %s, updated_at = NOW() WHERE id = :id", setQuery)
	fields["id"] = userUpdate.ID

	_, err := r.DB.NamedExec(query, fields)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresUserRepository) DeleteUser(id string) error {
	query := `UPDATE users SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		
		return err
	}
	return nil
}

func (r *PostgresUserRepository) GetUserInterviews(userID uuid.UUID) ([]models.Interview, error) {
	var interviews []models.Interview
	query := `SELECT * FROM interviews WHERE user_id = $1`
	err := r.DB.Select(&interviews, query, userID)
	if err != nil {
		return nil, err
	}
	return interviews, nil
}

func (r *PostgresUserRepository) GetUserResume(userID uuid.UUID) ([]models.Resume, error) {
	var resumes []models.Resume
	query := `SELECT * FROM resumes WHERE user_id = $1`
	err := r.DB.Select(&resumes, query, userID)
	if err != nil {
		return nil, err
	}
	return resumes, nil
}
