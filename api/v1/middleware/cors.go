package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware handles CORS headers for cross-origin requests.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Adjust the origins according to your frontend domain
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Change "*" to specific origin in production
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Handle preflight request
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
