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

	err = database.CreateUsersTable(db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	application := &app.Application{DB: db}
	studentRepository := repositories.NewSQLiteStudentRepository(application)
	studentService := services.NewStudentService(studentRepository)
	studentHandler := handlers.NewStudentHandler(studentService)

	routes.RegisterRoutes(studentHandler)

	fmt.Println("Server Running on :8080")

	http.ListenAndServe(":8080", nil)
}

// func add() { fmt.Println("Hello") }
