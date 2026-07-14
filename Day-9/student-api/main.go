package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var students = []Student{}

func StudentHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:

		json.NewEncoder(w).Encode(students)
	case http.MethodPost:
		var student Student
		err := json.NewDecoder(r.Body).Decode(&student)

		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if student.Name == "" {
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}

		if student.Age <= 0 {
			http.Error(w, "Age must be greater than 0", http.StatusBadRequest)
			return
		}

		students = append(students, student)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(student)

	default:

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {

	http.HandleFunc("/students", StudentHandler)

	fmt.Println("🚀 Server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server Error:", err)
	}
}