package database

import (
	"database/sql"
	"log"
	"time"
)

func New() *sql.DB {
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=student_db sslmode=disable"
	db, err := sql.Open(
		"postgres", //this just make connection pool
		dsn,
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err) // ping says hello postgresql
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	return db
}
