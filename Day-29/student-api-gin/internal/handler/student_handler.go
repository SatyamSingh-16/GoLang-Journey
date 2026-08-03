package handler

import (
	"net/http"
	"student-api-gin/internal/service"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	service *service.StudentService
}

func NewStudentHandler(
	service *service.StudentService,
) *StudentHandler {

	return &StudentHandler{
		service: service,
	}

}
func (h *StudentHandler) GetStudents(
	c *gin.Context,
) {
	ctx := c.Request.Context()
	students, err := h.service.GetStudents(ctx)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		students,
	)
}
