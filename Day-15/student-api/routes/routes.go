package routes

import (
	"net/http"

	"student-api/handlers"
)

func RegisterRoutes() {

	http.HandleFunc(
	"/students",

	Chain(

		handlers.StudentHandler,

		LoggingMiddleware,

		AuthMiddleware,

		ValidationMiddleware,

	),
)

}