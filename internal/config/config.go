package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
}

// LoadConfig загружает конфигурации из .env файла
func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	return Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:123456@localhost:5432/postgres?sslmode=disable"),
		ServerPort:  getEnv("SERVER_PORT", ":8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
