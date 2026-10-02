package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIKeyMiddleware(expectedKey string) gin.HandlerFunc {
	return func(c *gin.Context) {

		apiKey := c.GetHeader("X-API-Key")

		if apiKey != expectedKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthenticated access",
			})
			return
		}

		c.Next()
	}
}
