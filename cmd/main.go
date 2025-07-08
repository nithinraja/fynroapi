package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ai-financial-api/api/v1/middleware"
	"ai-financial-api/api/v1/router"
	"ai-financial-api/config"

	"github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, proceeding with environment variables.")
    }

    config.ConnectDatabase()

	r := middleware.CORS(router.SetupRouter())
	r = middleware.Logger(r)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8081"
    }

    fmt.Printf("Server running at http://localhost:%s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}
