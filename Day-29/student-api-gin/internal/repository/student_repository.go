package repository

import (
	"context"
	"database/sql"

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
