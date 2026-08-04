package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(c *gin.Context) {
	start := time.Now()
	method := c.Request.Method
	path := c.Request.URL.Path
	ip := c.ClientIP()
	c.Next()
	status := c.Writer.Status()
	duration := time.Since(start)
	fmt.Printf(
		"%s | %d | %v | %s | %s\n",
		method,
		status,
		duration,
		ip,
		path,
	)

}
