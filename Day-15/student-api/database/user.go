package database

import ("database/sql"
		"student-api/models")


func CreateUsersTable(db *sql.DB) error{
	query := `
	CREATE TABLE IF NOT EXISTS users(
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL
	)
	`

	_, err := db.Exec(query)
	return err;
}

func CreateUser(db *sql.DB, user models.User) error {
	query := `
		INSERT INTO users(
		id,
		name,
		email,
		password_hash
	)
	VALUES (?, ?, ?, ?)
	`

	_, err := db.Exec(
	query,
	user.ID,
	user.Name,
	user.Email,
	user.PasswordHash,
	)

return err
}

func GetUserByEmail(db *sql.DB, email string) (models.User, error) {
	query := `
		SELECT
		id,
		name,
		email,
		password_hash
	FROM users
	WHERE email = ?
	`
	row := db.QueryRow(query, email)

var user models.User

err := row.Scan(
	&user.ID,
	&user.Name,
	&user.Email,
	&user.PasswordHash,
)

return user, err
}