package v1

import (
	"github.com/gin-gonic/gin"

	"fyrnoapi/api/v1/handler"
	"fyrnoapi/api/v1/middleware"

	"github.com/sirupsen/logrus"
)

// RegisterV1Routes sets up all version 1 API routes and middleware
func RegisterV1Routes(router *gin.Engine, logger *logrus.Logger) {
	// Apply global middleware
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.LoggerMiddleware(logger))
	router.Use(gin.Recovery())

	// API V1 group
	v1 := router.Group("/api/v1")

	// Public Routes
	{
		v1.POST("/auth/request-otp", handler.RequestOTP)
		v1.POST("/auth/verify-otp", handler.VerifyOTP)

		v1.POST("/question", handler.AskQuestion) // with UUID
		v1.POST("/name", handler.SubmitName)      // assigns name to question and calls OpenAI

		v1.POST("/payment/initiate", handler.InitiatePayment)
		v1.POST("/payment/webhook", handler.PaymentWebhook)
	}

	// Authenticated Routes
	auth := v1.Group("/")
	auth.Use(middleware.JWTMiddleware())
	{
		auth.GET("/user/profile", handler.GetUserProfile)
		auth.GET("/user/questions", handler.GetUserQuestions)
		auth.POST("/user/update-name", handler.UpdateUserName)

		auth.POST("/question/analyze", handler.AnalyzeFinance)

		auth.GET("/recommendations", handler.GetRecommendations)
		auth.GET("/question/:id/recommendations", handler.GetRecommendationsByQuestionID)
	}
}
