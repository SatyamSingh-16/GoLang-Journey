package database

import (
	"database/sql"
	"fmt"
	"student-api/models"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {

	db, err := sql.Open(
		"sqlite",
		"student.db",
	)
	if err != nil {
		return nil, err
	}
	DB = db
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

func InsertStudent(db *sql.DB, student models.Student) (models.Student, error) {

	query := `
	INSERT INTO students (
		name,
		age
	)
	VALUES (?, ?)
	`

	result, err := db.Exec(
		query,
		student.Name,
		student.Age,
	)

	if err != nil {
		return student, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return student, err
	}

	student.ID = int(id)

	return student, nil
}

func GetStudents(db *sql.DB) ([]models.Student, error) {

	query := `
	SELECT
		id,
		name,
		age
	FROM students
	`

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var students []models.Student
	for rows.Next() {

		var student models.Student

		err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Age,
		)

		if err != nil {
			return nil, err
		}

		students = append(
			students,
			student,
		)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil

}

func GetStudentByID(db *sql.DB, id int) (models.Student, error) {

	query := `
	SELECT
		id,
		name,
		age
	FROM students
	WHERE id = ?
	`

	row := db.QueryRow(query, id)

	var student models.Student

	err := row.Scan(
		&student.ID,
		&student.Name,
		&student.Age,
	)

	if err != nil {
		return student, err
	}

	return student, nil
}
func UpdateStudent(db *sql.DB, student models.Student) error {

	query := `
	UPDATE students
	SET
		name = ?,
		age = ?
	WHERE id = ?
	`

	result, err := db.Exec(

		query,

		student.Name,

		student.Age,

		student.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func DeleteStudent(db *sql.DB, id int) error {

	query := `
	DELETE FROM students
	WHERE id = ?
	`

	result, err := db.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
