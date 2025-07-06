package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success sends a standard success response
func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

// Error sends a standard error response
func Error(c *gin.Context, statusCode int, message string, errors interface{}) {
	c.JSON(statusCode, gin.H{
		"success": false,
		"message": message,
		"errors":  errors,
	})
}

// Unauthorized sends a 401 response
func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, message, nil)
}

// BadRequest sends a 400 response
func BadRequest(c *gin.Context, message string, errors interface{}) {
	Error(c, http.StatusBadRequest, message, errors)
}

// NotFound sends a 404 response
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message, nil)
}

// InternalServerError sends a 500 response
func InternalServerError(c *gin.Context, message string, errors interface{}) {
	Error(c, http.StatusInternalServerError, message, errors)
}
