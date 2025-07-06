// =================
// cmd/server/main.go
// =================
package main

import (
	"log"
	"net/http"

	"fyrno.com/api/fyrnoapi/internal/config"
	"fyrno.com/api/fyrnoapi/internal/database"
	"fyrno.com/api/fyrnoapi/internal/handlers"
	"fyrno.com/api/fyrnoapi/internal/middleware"
	"fyrno.com/api/fyrnoapi/internal/services"
	"fyrno.com/api/fyrnoapi/pkg/utils"
	"github.com/gin-gonic/gin"
)

func main() {
    // Load configuration
    cfg := config.Load()

    // Initialize database
    if err := database.InitDB(cfg); err != nil {
        log.Fatal("Failed to initialize database:", err)
    }

    // Initialize services
    openaiService := services.NewOpenAIService(cfg.OpenAIAPIKey)
    questionService := services.NewQuestionService(openaiService)

    // Initialize handlers
    questionHandler := handlers.NewQuestionHandler(questionService)

    // Initialize Gin router
    router := gin.New()
    router.Use(middleware.Logger())
    router.Use(middleware.CORS())
    router.Use(gin.Recovery())

    // API routes
    api := router.Group("/api/v1")
    {
        // Question routes
        api.POST("/question", questionHandler.CreateQuestion)
        api.GET("/question/:questionid", questionHandler.GetQuestion)

        // Health check
        api.GET("/health", func(c *gin.Context) {
            utils.SuccessResponse(c, http.StatusOK, "Service is healthy", gin.H{
                "status": "healthy",
                "service": "FynroGo API",
            })
        })
    }

    // Start server
    log.Printf("Server starting on port %s", cfg.ServerPort)
    if err := router.Run(":" + cfg.ServerPort); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}