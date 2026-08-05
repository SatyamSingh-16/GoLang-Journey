package routes

import (
	"student-api-gin/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterStudentRoutes(
	router *gin.Engine,
	handler *handler.StudentHandler,
) {
	router.GET(
		"/students",
		handler.GetStudents,
	)
	router.GET(
		"/students/:id",
		handler.GetStudentByID,
	)
	router.POST(
		"/students",
		handler.CreateStudent,
	)
	router.PUT(
		"/students/:id",
		handler.UpdateStudent,
	)
}
