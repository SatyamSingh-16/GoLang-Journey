package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"student-api/models"
	"student-api/response"
)

func StudentHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:

		if r.URL.Path == "/students" {

			GetStudents(w)

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

func GetStudents(w http.ResponseWriter) {

	res := response.APIResponse{

		Success: true,
		Message: "Students fetched successfully",
		Data:    models.Students,
	}

	json.NewEncoder(w).Encode(res)

}

func CreateStudent(w http.ResponseWriter, r *http.Request) {

	var student models.Student

	json.NewDecoder(r.Body).Decode(&student)

	student.ID = len(models.Students) + 1

	models.Students = append(models.Students, student)

	res := response.APIResponse{

		Success: true,
		Message: "Student Created Successfully",
		Data:    student,
	}

	json.NewEncoder(w).Encode(res)

}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {

	id, err := GetIDFromURL(r)

	if err != nil {

		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return

	}

	for _, student := range models.Students {

		if student.ID == id {

			res := response.APIResponse{

				Success: true,
				Message: "Student Found",
				Data:    student,
			}

			json.NewEncoder(w).Encode(res)
			return

		}

	}

	http.Error(w, "Student Not Found", http.StatusNotFound)

}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	id, err := GetIDFromURL(r)

	if err != nil {

		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return

	}

	var updated models.Student

	json.NewDecoder(r.Body).Decode(&updated)

	for i := range models.Students {

		if models.Students[i].ID == id {

			updated.ID = id

			models.Students[i] = updated

			res := response.APIResponse{

				Success: true,
				Message: "Student Updated",
				Data:    updated,
			}

			json.NewEncoder(w).Encode(res)

			return

		}

	}

	http.Error(w, "Student Not Found", http.StatusNotFound)

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	id, err := GetIDFromURL(r)

	if err != nil {

		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return

	}

	for i := range models.Students {

		if models.Students[i].ID == id {

			models.Students = append(models.Students[:i], models.Students[i+1:]...)

			res := response.APIResponse{

				Success: true,
				Message: "Student Deleted",
				Data:    nil,
			}

			json.NewEncoder(w).Encode(res)

			return

		}

	}

	http.Error(w, "Student Not Found", http.StatusNotFound)

}

func GetIDFromURL(r *http.Request) (int, error) {

	parts := strings.Split(r.URL.Path, "/")

	return strconv.Atoi(parts[2])

}