package handler

import (
	"errors"
	"net/http"
	"student-api-gin/internal/dto"
	customerrors "student-api-gin/internal/errors"
	"student-api-gin/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(
	service *service.AuthService,
) *AuthHandler {

	return &AuthHandler{
		service: service,
	}
}
func (h *AuthHandler) Register(
	c *gin.Context,
) {
	ctx := c.Request.Context()
	var request dto.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	response, err := h.service.Register(
		ctx,
		request,
	)
	if errors.Is(err, customerrors.ErrEmailAlreadyExists) {
		c.JSON(
			http.StatusConflict,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	c.JSON(
		http.StatusInternalServerError,
		gin.H{
			"error": "internal server error",
		},
	)
	c.JSON(
		http.StatusCreated,
		response,
	)
}
func (h *AuthHandler) Login(
	c *gin.Context,
) {
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx := c.Request.Context()
	response, err := h.service.Login(
		ctx,
		request,
	)
	if err != nil {
		if errors.Is(
			err,
			customerrors.ErrInvalidCredentials,
		) {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "invalid credentials",
				},
			)
			return
		}
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
		response,
	)
}
