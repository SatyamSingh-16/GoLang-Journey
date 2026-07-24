package repositories

import (
	"database/sql"
	"student-api/app"
	apperrors "student-api/apperrors"
	"student-api/models"
)

type StudentRepository interface {
	GetStudents() ([]models.Student, error)
	CreateStudent(student models.Student) (models.Student, error)
	GetStudentByID(id int) (models.Student, error)
	UpdateStudent(student models.Student) error
	DeleteStudent(id int) error
}

type SQLiteStudentRepository struct {
	app *app.Application
}

func NewSQLiteStudentRepository(
	app *app.Application,
) *SQLiteStudentRepository {

	return &SQLiteStudentRepository{
		app: app,
	}
}

func (r *SQLiteStudentRepository) GetStudents() ([]models.Student, error) {
	query := `
	SELECT
		id,
		name,
		age
	FROM students
	`

	rows, err := r.app.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.ID, &student.Name, &student.Age); err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r *SQLiteStudentRepository) CreateStudent(student models.Student) (models.Student, error) {
	query := `
	INSERT INTO students(
		name,
		age
	)
	VALUES (?, ?)
	`

	result, err := r.app.DB.Exec(query, student.Name, student.Age)
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

func (r *SQLiteStudentRepository) GetStudentByID(id int) (models.Student, error) {
	query := `
	SELECT
		id,
		name,
		age
	FROM students
	WHERE id = ?
	`

	var student models.Student
	err := r.app.DB.QueryRow(query, id).Scan(&student.ID, &student.Name, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return student, apperrors.ErrStudentNotFound
		}
		return student, err
	}

	return student, nil
}

func (r *SQLiteStudentRepository) UpdateStudent(student models.Student) error {
	query := `
	UPDATE students
	SET
		name = ?,
		age = ?
	WHERE id = ?
	`

	result, err := r.app.DB.Exec(query, student.Name, student.Age, student.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.ErrStudentNotFound
	}

	return nil
}

func (r *SQLiteStudentRepository) DeleteStudent(id int) error {
	query := `
	DELETE FROM students
	WHERE id = ?
	`

	result, err := r.app.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.ErrStudentNotFound
	}

	return nil
}
