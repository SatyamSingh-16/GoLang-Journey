package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"student-api/database"
	"student-api/models"
	"student-api/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var req models.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w,
			"Invalid JSON",
			http.StatusBadRequest)
		return
	}
	if req.Name == "" {

		http.Error(
			w,
			"Name is required",
			http.StatusBadRequest,
		)

		return
	}
	if req.Email == "" {

		http.Error(
			w,
			"Email is required",
			http.StatusBadRequest,
		)

		return
	}
	if req.Password == "" {

		http.Error(
			w,
			"Password is required",
			http.StatusBadRequest,
		)

		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		http.Error(
			w,
			"Unable to hash password",
			http.StatusInternalServerError,
		)
		return
	}
	_, err = database.GetUserByEmail(database.DB, req.Email)

	if err == nil {

		http.Error(
			w,
			"Email already registered",
			http.StatusConflict,
		)

		return

	} else if err != sql.ErrNoRows {

		http.Error(
			w,
			"Database error",
			http.StatusInternalServerError,
		)

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
		http.Error(
			w,
			"Unable to create user",
			http.StatusInternalServerError,
		)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
	return
}

func Login(w http.ResponseWriter, r *http.Request) {

	var req models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {

		http.Error(
			w,
			"Invalid JSON",
			http.StatusBadRequest,
		)

		return
	}
	if req.Email == "" {

		http.Error(
			w,
			"Email is required",
			http.StatusBadRequest,
		)

		return
	}

	if req.Password == "" {

		http.Error(
			w,
			"Password is required",
			http.StatusBadRequest,
		)

		return
	}
	user, err := database.GetUserByEmail(
		database.DB,
		req.Email,
	)

	if err != nil {

		http.Error(
			w,
			"Invalid email or password",
			http.StatusUnauthorized,
		)

		return
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {

		http.Error(
			w,
			"Invalid email or password",
			http.StatusUnauthorized,
		)

		return
	}
	token, err := utils.GenerateToken(
		user.ID,
		user.Email,
	)

	if err != nil {

		http.Error(
			w,
			"Unable to generate token",
			http.StatusInternalServerError,
		)

		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})

}
