package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	apperrors "student-api/apperrors"
	"student-api/models"
	"student-api/response"
	"student-api/services"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	var req models.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	err = h.service.Register(req)
	if err != nil {
		if errors.Is(err, services.ErrUserNameRequired) ||
			errors.Is(err, services.ErrUserEmailRequired) ||
			errors.Is(err, services.ErrUserPasswordRequired) {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		if errors.Is(err, apperrors.ErrUserAlreadyExists) {
			response.Error(w, http.StatusConflict, "Email already registered")
			return
		}
		response.Error(w, http.StatusInternalServerError, "Unable to register user")
		return
	}

	response.Success(w, http.StatusCreated, "User registered successfully", nil)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	var req models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	token, err := h.service.Login(req)
	if err != nil {
		if errors.Is(err, services.ErrUserEmailRequired) ||
			errors.Is(err, services.ErrUserPasswordRequired) {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		if errors.Is(err, apperrors.ErrInvalidCredentials) {
			response.Error(w, http.StatusUnauthorized, "Invalid email or password")
			return
		}
		response.Error(w, http.StatusInternalServerError, "Unable to login")
		return
	}

	response.Success(w, http.StatusOK, "Login successful", map[string]string{
		"token": token,
	})
}
