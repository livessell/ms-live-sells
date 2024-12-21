package database

import (
	"fmt"
	"os"

	godotenv "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Carregar variáveis de ambiente do .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Não foi possível carregar o arquivo .env, usando variáveis do ambiente...")
	}

	var err error

	// Configurações de conexão
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "development" {
		DB, err = gorm.Open(postgres.Open("host=kesavan.db.elephantsql.com user=cfububgl password=hRL8JvCQoQ6Le6HDYGrc88jrOkGG00YY dbname=cfububgl sslmode=disable"), &gorm.Config{})
		if err != nil {
			panic("Falha ao conectar ao banco de testes!")
		}
	} else {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPassword, dbName,
		)

		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Falha ao conectar ao banco de dados!")
		}
	}
}
