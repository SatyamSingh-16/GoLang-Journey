package routes

import (
	"net/http"

	"student-api/handlers"
	"student-api/middleware"
)

func RegisterRoutes(studentHandler *handlers.StudentHandler) {

	http.HandleFunc(
		"/students",

		middleware.Chain(

			studentHandler.Handle,

			middleware.LoggingMiddleware,

			middleware.AuthMiddleware,

			middleware.ValidationMiddleware,
		),
	)

	http.HandleFunc(
		"/students/",
		middleware.Chain(
			studentHandler.Handle,
			middleware.LoggingMiddleware,
			middleware.AuthMiddleware,
			middleware.ValidationMiddleware,
		),
	)

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)

}
