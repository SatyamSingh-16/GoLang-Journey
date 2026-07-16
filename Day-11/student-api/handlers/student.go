package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"student-api/models"
)

func StudentHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		GetStudents(w, r)

	case http.MethodPost:
		CreateStudent(w, r)

	case http.MethodPut:
		UpdateStudent(w, r)

	case http.MethodPatch:
		PatchStudent(w, r)

	case http.MethodDelete:
		DeleteStudent(w, r)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func GetStudents(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(models.Students)

}

func CreateStudent(w http.ResponseWriter, r *http.Request) {

	var student models.Student

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	models.Students = append(models.Students, student)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(student)
}

func GetStudentByID(id int) *models.Student {

	for i := range models.Students {

		if models.Students[i].ID == id {

			return &models.Students[i]

		}
	}

	return nil
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	var updatedStudent models.Student

	err := json.NewDecoder(r.Body).Decode(&updatedStudent)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Temporary
	id, err := GetIDFromURL(r);
	student := GetStudentByID(id)

	if student == nil {
		http.Error(w, "Student Not Found", http.StatusNotFound)
		return
	}

	student.Name = updatedStudent.Name
	student.Age = updatedStudent.Age

	json.NewEncoder(w).Encode(student)
}

func PatchStudent(w http.ResponseWriter, r *http.Request) {

	var updatedStudent models.Student

	err := json.NewDecoder(r.Body).Decode(&updatedStudent)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Temporary
	id,err := GetIDFromURL(r);
	student := GetStudentByID(id)

	if student == nil {
		http.Error(w, "Student Not Found", http.StatusNotFound)
		return
	}

	if updatedStudent.Name != "" {
		student.Name = updatedStudent.Name
	}

	if updatedStudent.Age != 0 {
		student.Age = updatedStudent.Age
	}

	json.NewEncoder(w).Encode(student)
}

func DeleteStudentByID(id int) bool {

	for i := range models.Students {

		if models.Students[i].ID == id {

			models.Students = append(models.Students[:i], models.Students[i+1:]...)

			return true
		}
	}

	return false
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	// Temporary
	id,err := GetIDFromURL(r);
	deleted := DeleteStudentByID(id)

	if !deleted {
		http.Error(w, "Student Not Found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Student Deleted Successfully")
}

func GetIDFromURL(r *http.Request) (int, error) {

	parts := strings.Split(r.URL.Path, "/")

	id, err := strconv.Atoi(parts[2])

	if err != nil {

		return 0, err

	}

	return id, nil
}