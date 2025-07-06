package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggerMiddleware logs HTTP requests using logrus with detailed information.
func LoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start time
		startTime := time.Now()

		// Process request
		c.Next()

		// End time
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// Get status
		statusCode := c.Writer.Status()

		// Log fields
		logger.WithFields(logrus.Fields{
			"status":   statusCode,
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
			"latency":  latency,
			"ip":       c.ClientIP(),
			"userAgent": c.Request.UserAgent(),
		}).Info("HTTP Request")
	}
}
