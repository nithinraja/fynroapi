// =================
// internal/config/config.go
// =================
package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    DatabaseHost     string
    DatabasePort     string
    DatabaseUser     string
    DatabasePassword string
    DatabaseName     string
    ServerPort       string
    OpenAIAPIKey     string
}

func Load() *Config {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using environment variables")
    }

    return &Config{
        DatabaseHost:     getEnv("DB_HOST", "localhost"),
        DatabasePort:     getEnv("DB_PORT", "3306"),
        DatabaseUser:     getEnv("DB_USER", "root"),
        DatabasePassword: getEnv("DB_PASSWORD", ""),
        DatabaseName:     getEnv("DB_NAME", "fynroapi"),
        ServerPort:       getEnv("SERVER_PORT", "8080"),
        OpenAIAPIKey:     getEnv("OPENAI_API_KEY", ""),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}