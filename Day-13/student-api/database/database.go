package database

var DB *sql.DB
import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func ConnectDB() (*sql.DB, error) {
	DB = db

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

func InsertStudent(student models.Student) (models.Student, error) {

	query := `
	INSERT INTO students (
		name,
		age
	)
	VALUES (?, ?)
	`

	result, err := DB.Exec(
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

func GetStudents() ([]models.Student, error) {

	query := `
	SELECT
		id,
		name,
		age
	FROM students
	`

	rows, err := DB.Query(query)

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