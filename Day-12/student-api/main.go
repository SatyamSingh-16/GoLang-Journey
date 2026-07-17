package main

import (
	"fmt"
	"net/http"

	"student-api/handlers"
	"student-api/middleware"
)

func main() {

	http.HandleFunc(

		"/students",

		middleware.Chain(

			handlers.StudentHandler,

			middleware.LoggingMiddleware,

			middleware.AuthMiddleware,

			middleware.ValidationMiddleware,
		),
	)

	http.HandleFunc(

		"/students/",

		middleware.Chain(

			handlers.StudentHandler,

			middleware.LoggingMiddleware,

			middleware.AuthMiddleware,

			middleware.ValidationMiddleware,
		),
	)

	fmt.Println("Server Running on :8080")

	http.ListenAndServe(":8080", nil)

}