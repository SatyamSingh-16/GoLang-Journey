package main

import (
	"fmt"
	"net/http"

	"student-api/app"
	"student-api/config"
	"student-api/database"
	"student-api/handlers"
	"student-api/repositories"
	"student-api/routes"
	"student-api/services"
)

func main() {
	config.LoadConfig()

	// Connect to Database
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Initialize Database Schema
	if err := database.InitSchema(db); err != nil {
		panic(err)
	}

	application := &app.Application{DB: db}

	// Repositories
	studentRepository := repositories.NewPostgresStudentRepository(application)
	userRepository := repositories.NewPostgresUserRepository(application)


	// Services
	studentService := services.NewStudentService(studentRepository)
	userService := services.NewUserService(userRepository)

	// Handlers
	studentHandler := handlers.NewStudentHandler(studentService)
	userHandler := handlers.NewUserHandler(userService)

	// Register Routes
	routes.RegisterRoutes(studentHandler, userHandler)

	fmt.Println("Server Running on :8080")
	http.ListenAndServe(":8080", nil)
}

