package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load environment variables from .env file
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or could not be loaded.")
	}
}

// Get environment variable or log fatal error if not set
func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s environment variable not set", key)
	}
	return value
}
