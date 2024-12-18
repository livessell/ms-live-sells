package config

import (
	"github.com/joho/godotenv"
	"log"
)

func Init() {
	// Carrega o arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Nenhum arquivo .env encontrado, usando variáveis de ambiente do sistema.")
	}
}
