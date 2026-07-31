package repositories

import (
	"student-api/app"
	"student-api/models"
)

type UserRepository interface {
	CreateUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
}

type PostgresUserRepository struct {
	app *app.Application
}

func NewPostgresUserRepository(app *app.Application) *PostgresUserRepository {
	return &PostgresUserRepository{
		app: app,
	}
}

// Keeping NewSQLiteUserRepository for backward compatibility
func NewSQLiteUserRepository(app *app.Application) *PostgresUserRepository {
	return NewPostgresUserRepository(app)
}

func (r *PostgresUserRepository) CreateUser(user models.User) error {
	query := `
	INSERT INTO users(
		id,
		name,
		email,
		password_hash
	)
	VALUES ($1, $2, $3, $4)
	`

	_, err := r.app.DB.Exec(
		query,
		user.ID,
		user.Name,
		user.Email,
		user.PasswordHash,
	)

	return err
}

func (r *PostgresUserRepository) GetUserByEmail(email string) (models.User, error) {
	query := `
	SELECT
		id,
		name,
		email,
		password_hash
	FROM users
	WHERE email = $1
	`

	var user models.User
	err := r.app.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
	)

	return user, err
}
