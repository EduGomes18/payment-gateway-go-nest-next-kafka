package main

import (
	"database/sql"
	"fmt"
	"go-gateway-api/internal/repository"
	"go-gateway-api/internal/service"
	"go-gateway-api/internal/web/server"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "postgres"),
		getEnv("DB_SSL_MODE", "disable"),
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()


	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	port := getEnv("HTTP_PORT", "8080")
	
	srv := server.NewServer(port, accountService)
	srv.ConfigureRoutes()
	
	if err := srv.Start(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

