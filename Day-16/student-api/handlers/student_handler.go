package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"student-api/database"
	"student-api/models"
	"student-api/response"
)

func StudentHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:

		if r.URL.Path == "/students" {

			GetStudents(w, r)

		} else {

			GetStudentByID(w, r)

		}

	case http.MethodPost:

		CreateStudent(w, r)

	case http.MethodPut:

		UpdateStudent(w, r)

	case http.MethodDelete:

		DeleteStudent(w, r)

	default:

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

	}

}

func GetStudents(w http.ResponseWriter, r *http.Request) {

	students, err := database.GetStudents()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	res := response.APIResponse{

		Success: true,

		Message: "Students fetched successfully",

		Data: students,
	}

	json.NewEncoder(w).Encode(res)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {

	var student models.Student

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {

		http.Error(
			w,
			"Invalid JSON",
			http.StatusBadRequest,
		)
		return
	}

	student,err := database.InsertStudent(student)
	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	res := response.APIResponse{
		Success: true,
		Message: "Student Created Successfully",
		Data:    student, // ID will be fixed in the next lesson
	}

	json.NewEncoder(w).Encode(res)
}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(
		r.URL.Path,
		"/students/",
	)

	id, err := strconv.Atoi(idStr)
	if err != nil {

		http.Error(
			w,
			"Invalid Student ID",
			http.StatusBadRequest,
		)

		return
	}

	student, err := database.GetStudentByID(id)

	if err != nil {

		if err == sql.ErrNoRows {

			http.Error(
				w,
				"Student Not Found",
				http.StatusNotFound,
			)

			return
		}

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	res := response.APIResponse{

		Success: true,

		Message: "Student fetched successfully",

		Data: student,
	}

	json.NewEncoder(w).Encode(res)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(
		r.URL.Path,
		"/students/",
	)

	id, err := strconv.Atoi(idStr)
	if err != nil {

		http.Error(
			w,
			"Invalid Student ID",
			http.StatusBadRequest,
		)

		return
	}

	var student models.Student

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {

		http.Error(
			w,
			"Invalid JSON",
			http.StatusBadRequest,
		)

		return
	}

	student.ID = id

	err = database.UpdateStudent(student)

	if err != nil {

		if err == sql.ErrNoRows {

			http.Error(
				w,
				"Student Not Found",
				http.StatusNotFound,
			)

			return
		}

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	res := response.APIResponse{

		Success: true,

		Message: "Student Updated Successfully",

		Data: student,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(
		r.URL.Path,
		"/students/",
	)

	id, err := strconv.Atoi(idStr)

	if err != nil {

		http.Error(
			w,
			"Invalid Student ID",
			http.StatusBadRequest,
		)

		return
	}

	err = database.DeleteStudent(id)

	if err != nil {

		if err == sql.ErrNoRows {

			http.Error(
				w,
				"Student Not Found",
				http.StatusNotFound,
			)

			return
		}

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	res := response.APIResponse{

		Success: true,

		Message: "Student Deleted Successfully",

		Data: nil,
	}

	json.NewEncoder(w).Encode(res)
}

func GetIDFromURL(r *http.Request) (int, error) {

	parts := strings.Split(r.URL.Path, "/")

	return strconv.Atoi(parts[2])

}