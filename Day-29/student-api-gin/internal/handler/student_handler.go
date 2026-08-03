package handler

import (
	"database/sql"
	"net/http"
	"strconv"
	"student-api-gin/internal/dto"
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

func (h *StudentHandler) GetStudentByID(
	c *gin.Context,
) {
	ctx := c.Request.Context()
	id := c.Param("id")
	studentID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid student id",
			},
		)
		return
	}
	student, err := h.service.GetStudentByID(
		ctx,
		studentID,
	)
	if err == sql.ErrNoRows {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "student not found",
			},
		)
		return
	}
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "internal server error",
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		student,
	)
}
func (h *StudentHandler) CreateStudent(
	c *gin.Context,
) {

	ctx := c.Request.Context()

	var request dto.CreateStudentRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	id, err := h.service.CreateStudent(
		ctx,
		request,
	)
	c.JSON(
		http.StatusCreated,
		gin.H{
			"id":      id,
			"message": "student created successfully",
		},
	)

}
