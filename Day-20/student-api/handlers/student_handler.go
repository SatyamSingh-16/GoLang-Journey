package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	apperrors "student-api/apperrors"
	"student-api/models"
	"student-api/response"
	"student-api/services"
)

type StudentHandler struct {
	service *services.StudentService
}

func NewStudentHandler(service *services.StudentService) *StudentHandler {

	return &StudentHandler{
		service: service,
	}
}

func (h *StudentHandler) Handle(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:

		if r.URL.Path == "/students" {

			h.GetStudents(w, r)

		} else {

			h.GetStudentByID(w, r)

		}

	case http.MethodPost:

		h.CreateStudent(w, r)

	case http.MethodPut:

		h.UpdateStudent(w, r)

	case http.MethodDelete:

		h.DeleteStudent(w, r)

	default:

		response.Error(w, http.StatusMethodNotAllowed, "Method Not Allowed")

	}

}

func (h *StudentHandler) GetStudents(w http.ResponseWriter, r *http.Request) {

	students, err := h.service.GetStudents()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "Students fetched successfully", students)
}

func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {

	var student models.Student

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {

		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	student, err := h.service.CreateStudent(student)
	if err != nil {
		if errors.Is(err, services.ErrStudentNameRequired) || errors.Is(err, services.ErrStudentAgeInvalid) {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusCreated, "Student Created Successfully", student)
}

func (h *StudentHandler) GetStudentByID(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(
		r.URL.Path,
		"/students/",
	)

	id, err := strconv.Atoi(idStr)
	if err != nil {

		response.Error(w, http.StatusBadRequest, "Invalid Student ID")

		return
	}

	student, err := h.service.GetStudentByID(id)

	if err != nil {

		if errors.Is(err, apperrors.ErrStudentNotFound) {

			response.Error(w, http.StatusNotFound, "Student Not Found")

			return
		}

		response.Error(w, http.StatusInternalServerError, err.Error())

		return
	}

	response.Success(w, http.StatusOK, "Student fetched successfully", student)
}

func (h *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {

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

	err = h.service.UpdateStudent(student)

	if err != nil {
		if errors.Is(err, services.ErrStudentNameRequired) || errors.Is(err, services.ErrStudentAgeInvalid) {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		if errors.Is(err, apperrors.ErrStudentNotFound) {
			response.Error(w, http.StatusNotFound, "Student Not Found")
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "Student Updated Successfully", student)
}

func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(
		r.URL.Path,
		"/students/",
	)

	id, err := strconv.Atoi(idStr)

	if err != nil {

		response.Error(w, http.StatusBadRequest, "Invalid Student ID")

		return
	}

	err = h.service.DeleteStudent(id)

	if err != nil {
		if errors.Is(err, apperrors.ErrStudentNotFound) {
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
