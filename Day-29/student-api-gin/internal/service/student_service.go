package service

import (
	"context"
	"student-api-gin/internal/models"
	"student-api-gin/internal/repository"
)

type StudentService struct {
	repo *repository.StudentRepository
}

func NewStudentService(
	repo *repository.StudentRepository,
) *StudentService {

	return &StudentService{
		repo: repo,
	}

}
func (s *StudentService) GetStudents(
	ctx context.Context,
) ([]models.Student, error) {

	return s.repo.GetStudents(ctx)

}
