package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() error {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return err
	}
	return nil
}

func Get(key string, defaultValue string) string {
	if defaultValue == "" {
		return os.Getenv(key)
	}
	return defaultValue
}
