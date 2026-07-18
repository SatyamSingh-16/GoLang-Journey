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

	fmt.Println("Database Connected Successfully ✅")

	return db, nil
}