package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB() (*sql.DB, error) {
	connectionString :=
		"host=localhost user=satyamsingh2730 dbname=studentdb sslmode=disable"

	db, err := sql.Open(
		"pgx",
		connectionString,
	)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Database Connected Successfully ✅")
	return db, nil
}

func InitSchema(db *sql.DB) error {
	if err := CreateStudentsTable(db); err != nil {
		return err
	}
	return CreateUsersTable(db)
}

func CreateStudentsTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS students (
			id SERIAL PRIMARY KEY ,
			name TEXT NOT NULL,
			age INTEGER NOT NULL
		);`
	_, err := db.Exec(query)
	return err
}
