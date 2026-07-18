package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func ConnectDB() (*sql.DB, error) {

	db, err := sql.Open(
		"sqlite",
		"student.db",
	)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	query := `
		CREATE TABLE IF NOT EXISTS students (

		id INTEGER PRIMARY KEY AUTOINCREMENT,

		name TEXT NOT NULL,

		age INTEGER NOT NULL

	);`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	fmt.Println("Database Connected Successfully ✅")

	return db, nil
}