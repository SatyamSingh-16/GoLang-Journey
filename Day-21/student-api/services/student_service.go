package services

import (
	"errors"
	"student-api/models"
	"student-api/repositories"
)

type StudentService struct {
	repository repositories.StudentRepository
}

var (
	ErrStudentNameRequired = errors.New("student name is required")
	ErrStudentAgeInvalid   = errors.New("student age must be greater than zero")
)

func NewStudentService(repository repositories.StudentRepository) *StudentService {

	return &StudentService{
		repository: repository,
	}
}

func (s *StudentService) CreateStudent(
	student models.Student,
) (models.Student, error) {
	if err := validateStudent(student); err != nil {
		return student, err
	}

	return s.repository.CreateStudent(student)
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.repository.GetStudents()
}

func (s *StudentService) GetStudentByID(id int) (models.Student, error) {
	return s.repository.GetStudentByID(id)
}

func (s *StudentService) UpdateStudent(student models.Student) error {
	if err := validateStudent(student); err != nil {
		return err
	}

	return s.repository.UpdateStudent(student)
}

func (s *StudentService) DeleteStudent(id int) error {
	return s.repository.DeleteStudent(id)
}

func validateStudent(student models.Student) error {
	if student.Name == "" {
		return ErrStudentNameRequired
	}

	if student.Age <= 0 {
		return ErrStudentAgeInvalid
	}

	return nil
}
