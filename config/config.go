package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Getenv(key string) string {
	return os.Getenv(key)
}

func Loadenv() {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}