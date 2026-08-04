package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger(c *gin.Context) {
	fmt.Println("Incoming Request")
	c.Next()
}
