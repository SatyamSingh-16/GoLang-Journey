package main

import (
	"student-api-gin/internal/database"
	"student-api-gin/internal/handler"
	"student-api-gin/internal/middleware"
	"student-api-gin/internal/repository"
	"student-api-gin/internal/routes"
	"student-api-gin/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(middleware.Logger)
	db := database.New()
	repo := repository.NewStudentRepository(db)
	service := service.NewStudentService(repo)
	handler := handler.NewStudentHandler(service)
	routes.RegisterStudentRoutes(
		router,
		handler,
	)
	router.Run(":8080")

}
