package services

import (
	"database/sql"
	"errors"
	apperrors "student-api/apperrors"
	"student-api/models"
	"student-api/repositories"
	"student-api/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNameRequired     = errors.New("name is required")
	ErrUserEmailRequired    = errors.New("email is required")
	ErrUserPasswordRequired = errors.New("password is required")
)

type UserService interface {
	Register(req models.RegisterRequest) error
	Login(req models.LoginRequest) (string, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Register(req models.RegisterRequest) error {
	if req.Name == "" {
		return ErrUserNameRequired
	}
	if req.Email == "" {
		return ErrUserEmailRequired
	}
	if req.Password == "" {
		return ErrUserPasswordRequired
	}

	_, err := s.userRepo.GetUserByEmail(req.Email)
	if err == nil {
		return apperrors.ErrUserAlreadyExists
	} else if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user := models.User{
		ID:           uuid.New().String(),
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	return s.userRepo.CreateUser(user)
}

func (s *userService) Login(req models.LoginRequest) (string, error) {
	if req.Email == "" {
		return "", ErrUserEmailRequired
	}
	if req.Password == "" {
		return "", ErrUserPasswordRequired
	}

	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", apperrors.ErrInvalidCredentials
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)
	if err != nil {
		return "", apperrors.ErrInvalidCredentials
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
