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

		response.Error(w, http.StatusMethodNotAllowed, "Method Not Allowed")

	}

}

func GetStudents(w http.ResponseWriter, r *http.Request) {

	students, err := database.GetStudents()

	if err != nil {

		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	response.Success(w, http.StatusOK, "Students fetched successfully", students)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {

	var student models.Student

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {

		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	student, err := database.InsertStudent(student)
	if err != nil {

		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusCreated, "Student Created Successfully", student)
}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(
		r.URL.Path,
		"/students/",
	)

	id, err := strconv.Atoi(idStr)
	if err != nil {

		response.Error(w, http.StatusBadRequest, "Invalid Student ID")

		return
	}

	student, err := database.GetStudentByID(id)

	if err != nil {

		if err == sql.ErrNoRows {

			response.Error(w, http.StatusNotFound, "Student Not Found")

			return
		}

		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	response.Success(w, http.StatusOK, "Student fetched successfully", student)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(
		r.URL.Path,
		"/students/",
	)

	id, err := strconv.Atoi(idStr)
	if err != nil {

		response.Error(w, http.StatusBadRequest, "Invalid Student ID")

		return
	}

	var student models.Student

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {

		response.Error(w, http.StatusBadRequest, "Invalid JSON")

		return
	}

	student.ID = id

	err = database.UpdateStudent(student)

	if err != nil {

		if err == sql.ErrNoRows {

			response.Error(w, http.StatusNotFound, "Student Not Found")

			return
		}

		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	response.Success(w, http.StatusOK, "Student Updated Successfully", student)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(
		r.URL.Path,
		"/students/",
	)

	id, err := strconv.Atoi(idStr)

	if err != nil {

		response.Error(w, http.StatusBadRequest, "Invalid Student ID")

		return
	}

	err = database.DeleteStudent(id)

	if err != nil {

		if err == sql.ErrNoRows {

			response.Error(w, http.StatusNotFound, "Student Not Found")

			return
		}

		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	response.Success(w, http.StatusOK, "Student Deleted Successfully", nil)
}

func GetIDFromURL(r *http.Request) (int, error) {

	parts := strings.Split(r.URL.Path, "/")

	return strconv.Atoi(parts[2])

}
