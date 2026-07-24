package routes

import (
	"net/http"

	"student-api/handlers"
	"student-api/middleware"
)

func RegisterRoutes(studentHandler *handlers.StudentHandler, userHandler *handlers.UserHandler) {

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

	http.HandleFunc("/register", userHandler.Register)
	http.HandleFunc("/login", userHandler.Login)

}

