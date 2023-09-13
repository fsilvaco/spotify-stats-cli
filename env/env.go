package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get(key string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return os.Getenv(key)
}
