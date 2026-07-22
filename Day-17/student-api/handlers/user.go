package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"student-api/database"
	"student-api/models"
	"student-api/response"
	"student-api/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var req models.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	if req.Name == "" {

		response.Error(w, http.StatusBadRequest, "Name is required")

		return
	}
	if req.Email == "" {

		response.Error(w, http.StatusBadRequest, "Email is required")

		return
	}
	if req.Password == "" {

		response.Error(w, http.StatusBadRequest, "Password is required")

		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Unable to hash password")
		return
	}
	_, err = database.GetUserByEmail(database.DB, req.Email)

	if err == nil {

		response.Error(w, http.StatusConflict, "Email already registered")

		return

	} else if err != sql.ErrNoRows {

		response.Error(w, http.StatusInternalServerError, "Database error")

		return
	}
	user := models.User{
		ID:           uuid.New().String(),
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	err = database.CreateUser(database.DB, user)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Unable to create user")
		return
	}

	response.Success(w, http.StatusCreated, "User registered successfully", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {

	var req models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {

		response.Error(w, http.StatusBadRequest, "Invalid JSON")

		return
	}
	if req.Email == "" {

		response.Error(w, http.StatusBadRequest, "Email is required")

		return
	}

	if req.Password == "" {

		response.Error(w, http.StatusBadRequest, "Password is required")

		return
	}
	user, err := database.GetUserByEmail(
		database.DB,
		req.Email,
	)

	if err != nil {

		response.Error(w, http.StatusUnauthorized, "Invalid email or password")

		return
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {

		response.Error(w, http.StatusUnauthorized, "Invalid email or password")

		return
	}
	token, err := utils.GenerateToken(
		user.ID,
		user.Email,
	)

	if err != nil {

		response.Error(w, http.StatusInternalServerError, "Unable to generate token")

		return
	}
	response.Success(w, http.StatusOK, "Login successful", map[string]string{
		"token": token,
	})

}
