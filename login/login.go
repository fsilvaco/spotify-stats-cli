package login

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Inicialize() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
	port := os.Getenv("PORT")
	fmt.Println("Here the port is", port)
}
