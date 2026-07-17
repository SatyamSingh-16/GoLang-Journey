package routes

import (
	"net/http"

	"student-api/handlers"
)

func RegisterRoutes() {

	http.HandleFunc("/students", handlers.StudentHandler)

}