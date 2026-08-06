package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(
	allowedRoles ...string,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		roleValue, exists := c.Get("role")

		if !exists {

			c.JSON(http.StatusForbidden, gin.H{
				"error": "role not found",
			})

			c.Abort()

			return
		}

		role := roleValue.(string)

		authorized := false

		for _, allowed := range allowedRoles {

			if role == allowed {

				authorized = true

				break

			}

		}

		if !authorized {

			c.JSON(http.StatusForbidden, gin.H{
				"error": "access denied",
			})

			c.Abort()

			return
		}

		c.Next()
	}

}
