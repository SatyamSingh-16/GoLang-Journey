package routes

import (
	"student-api-gin/internal/handler"
	"student-api-gin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterStudentRoutes(
	router *gin.Engine,
	handler *handler.StudentHandler,
) {
	protected := router.Group("/")

	protected.Use(
		middleware.AuthMiddleware(),
	)

	protected.GET(
		"/students",
		handler.GetStudents,
	)

	protected.POST(
		"/students",
		handler.CreateStudent,
	)
	protected.GET(
		"/students/:id",
		handler.GetStudentByID,
	)

	protected.PUT(
		"/students/:id",
		handler.UpdateStudent,
	)
}
