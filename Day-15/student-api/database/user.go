package database

import "database/sql"


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