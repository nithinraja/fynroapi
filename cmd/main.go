package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	v1 "fyrnoapi/api/v1"
	"fyrnoapi/config"
	"fyrnoapi/pkg/database"
	"fyrnoapi/pkg/logger"
)

func main() {
	// Load .env and configuration
	config.LoadConfig()

	// Initialize logger
	log := logger.NewLogger()
	log.Info("Starting Finance AI backend...")

	// Connect to database
	err := database.ConnectDB()
	if err != nil {
		log.WithError(err).Fatal("Could not connect to the database")
		os.Exit(1)
	}

	// Run database migrations
	err = database.RunMigrations()
	if err != nil {
		log.WithError(err).Fatal("Could not run migrations")
		os.Exit(1)
	}

	// Initialize Gin router
	router := gin.New()

	// Register API routes
	v1.RegisterV1Routes(router, log)

	// Start server
	port := config.AppConfig.ServerPort
	log.Infof("Server is running at http://localhost:%s", port)
	err = router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.WithError(err).Fatal("Failed to run server")
	}
}
