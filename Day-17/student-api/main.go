package main

import (
	"fmt"
	"net/http"

	"student-api/app"
	"student-api/config"
	"student-api/database"
	"student-api/handlers"
	"student-api/middleware"
)

func main() {
	application := &app.Application{
		DB: database.DB,
	}
	config.LoadConfig()
	// Connect to Database
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = database.CreateUsersTable(db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Register Routes
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

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)

	fmt.Println("Server Running on :8080")

	http.ListenAndServe(":8080", nil)
}

// func add() { fmt.Println("Hello") }
