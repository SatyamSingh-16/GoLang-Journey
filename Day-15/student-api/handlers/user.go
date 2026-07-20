package handlers

import (
	"encoding/json"
	"net/http"
	"student-api/database"
	"student-api/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var req models.RegisterRequest

	err := json.NewDecoder(r.body).Decode(&req)
	if err != nil{
		htt.Error(w,
		"Invalid JSON",
		http.StatusBadRequest,)
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