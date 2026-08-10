package repository

import (
	"context"
	"database/sql"
	"student-api-gin/internal/models"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

// CreateUser inserts a new user into the database.
func (r *AuthRepository) CreateUser(
	ctx context.Context,
	user *models.User,
) error {

	query := `
		INSERT INTO users
		(name, email, password_hash, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at;
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.Role,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

// GetUserByEmail returns a user using email.
func (r *AuthRepository) GetUserByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {

	query := `
		SELECT
			id,
			name,
			email,
			password_hash,
			role,
			created_at
		FROM users
		WHERE email = $1;
	`

	var user models.User

	err := r.db.QueryRowContext(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByID returns a user using id.
func (r *AuthRepository) GetUserByID(
	ctx context.Context,
	id int,
) (*models.User, error) {

	query := `
		SELECT
			id,
			name,
			email,
			password_hash,
			role,
			created_at
		FROM users
		WHERE id = $1;
	`

	var user models.User

	err := r.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
