package middleware

import (
	"net/http"
	"strings"
	"student-api-gin/internal/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {

			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header missing",
			})

			c.Abort()

			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 ||
			parts[0] != "Bearer" {

			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header",
			})

			c.Abort()

			return
		}

		claims, err := auth.ValidateJWT(parts[1])

		if err != nil {

			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})

			c.Abort()

			return
		}

		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}
