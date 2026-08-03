package repository

import (
	"context"
	"database/sql"

	"student-api-gin/internal/dto"
	"student-api-gin/internal/models"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

func (r *StudentRepository) GetStudents(
	ctx context.Context,
) ([]models.Student, error) {
	query := `
	SELECT
		id,
		name,
		email,
		age
	FROM students
`
	rows, err := r.db.QueryContext(
		ctx,
		query,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	students := []models.Student{}

	for rows.Next() {

		var student models.Student
		err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Email,
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
func (r *StudentRepository) GetStudentByID(
	ctx context.Context,
	id int,
) (*models.Student, error) {
	query := `
	SELECT id, name, email
	FROM students
	WHERE id = $1
`
	row := r.db.QueryRowContext(
		ctx,
		query,
		id,
	)
	var student models.Student
	err := row.Scan(
		&student.ID,
		&student.Name,
		&student.Email,
	)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &student, nil
}
func (r *StudentRepository) CreateStudent(
	ctx context.Context,
	request dto.CreateStudentRequest,
) (int, error) {
	query := `
	INSERT INTO students(name,email)
	VALUES($1,$2)
	RETURNING id
	`
	var id int

	err := r.db.QueryRowContext(
		ctx,
		query,
		request.Name,
		request.Email,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
