package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all the configuration values loaded from .env or environment
type Config struct {
	AppEnv           string
	ServerPort       string
	DBUser           string
	DBPassword       string
	DBHost           string
	DBPort           string
	DBName           string
	JWTSecret        string
	TwilioAccountSID string
	TwilioAuthToken  string
	TwilioFromNumber string
	OpenAIAPIKey     string
	RazorpayKey      string
	RazorpaySecret   string
}

// AppConfig is the global configuration object
var AppConfig *Config

// LoadConfig initializes configuration from .env or system environment
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	AppConfig = &Config{
		AppEnv:           getEnv("APP_ENV", "development"),
		ServerPort:       getEnv("SERVER_PORT", "8080"),
		DBUser:           getEnv("DB_USER", "root"),
		DBPassword:       getEnv("DB_PASSWORD", ""),
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "3306"),
		DBName:           getEnv("DB_NAME", "finance_ai"),
		JWTSecret:        getEnv("JWT_SECRET", "supersecretjwtkey"),
		TwilioAccountSID: getEnv("TWILIO_ACCOUNT_SID", ""),
		TwilioAuthToken:  getEnv("TWILIO_AUTH_TOKEN", ""),
		TwilioFromNumber: getEnv("TWILIO_FROM_NUMBER", ""),
		OpenAIAPIKey:     getEnv("OPENAI_API_KEY", ""),
		RazorpayKey:      getEnv("RAZORPAY_KEY", ""),
		RazorpaySecret:   getEnv("RAZORPAY_SECRET", ""),
	}
}

// getEnv retrieves an environment variable or returns a fallback
func getEnv(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}
