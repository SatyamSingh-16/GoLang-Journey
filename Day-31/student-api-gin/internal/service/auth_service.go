package service

import (
	"context"
	"database/sql"
	"errors"
	"student-api-gin/internal/dto"
	customerrors "student-api-gin/internal/errors"
	"student-api-gin/internal/models"
	"student-api-gin/internal/repository"

	"student-api-gin/internal/auth"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(
	repo *repository.AuthRepository,
) *AuthService {

	return &AuthService{
		repo: repo,
	}
}
func (s *AuthService) Register(
	ctx context.Context,
	request dto.RegisterRequest,
) (*dto.RegisterResponse, error) {
	existingUser, err := s.repo.GetUserByEmail(
		ctx,
		request.Email,
	)

	if err != nil {

		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

	}

	if existingUser != nil {
		return nil, customerrors.ErrEmailAlreadyExists
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}
	user := models.User{
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: string(hashedPassword),
		Role:         "student",
	}
	err = s.repo.CreateUser(
		ctx,
		&user,
	)

	if err != nil {
		return nil, err
	}
	response := dto.RegisterResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Message: "User registered successfully",
	}
	return &response, nil
}
func (s *AuthService) Login(
	ctx context.Context,
	request dto.LoginRequest,
) (*dto.LoginResponse, error) {

	user, err := s.repo.GetUserByEmail(
		ctx,
		request.Email,
	)

	if err != nil {
		return nil, customerrors.ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(request.Password),
	)

	if err != nil {
		return nil, customerrors.ErrInvalidCredentials
	}

	token, err := auth.GenerateJWT(user)

	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: token,
	}, nil
}
