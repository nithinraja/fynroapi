package middleware

import (
	"net/http"
	"strings"

	"fyrnoapi/pkg/token"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware checks for a valid JWT token in the Authorization header.
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := token.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Store user ID in context
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
