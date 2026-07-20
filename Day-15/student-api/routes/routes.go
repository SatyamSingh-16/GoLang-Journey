package routes

import (
	"net/http"

	"student-api/handlers"
	"student-api/middleware"
)

func RegisterRoutes() {

	http.HandleFunc(
	"/students",

	middleware.Chain(

		handlers.StudentHandler,

		middleware.LoggingMiddleware,

		middleware.AuthMiddleware,

		middleware.ValidationMiddleware,

	),
)

}