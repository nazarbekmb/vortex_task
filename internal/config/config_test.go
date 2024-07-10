package config_test

import (
	"os"
	"testing"

	"statistics-collection/internal/config"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgres://user:password@localhost:5432/testdb?sslmode=disable")
	os.Setenv("SERVER_PORT", ":8081")

	cfg := config.LoadConfig()

	expectedDBURL := "postgres://user:password@localhost:5432/testdb?sslmode=disable"
	if cfg.DatabaseURL != expectedDBURL {
		t.Errorf("Expected DatabaseURL %s, got %s", expectedDBURL, cfg.DatabaseURL)
	}

	expectedServerPort := ":8081"
	if cfg.ServerPort != expectedServerPort {
		t.Errorf("Expected ServerPort %s, got %s", expectedServerPort, cfg.ServerPort)
	}
}

func TestLoadConfig_DefaultValues(t *testing.T) {
	os.Unsetenv("DATABASE_URL")
	os.Unsetenv("SERVER_PORT")

	cfg := config.LoadConfig()

	expectedDBURL := "postgres://postgres:123456@localhost:5432/postgres?sslmode=disable"
	if cfg.DatabaseURL != expectedDBURL {
		t.Errorf("Expected DatabaseURL %s, got %s", expectedDBURL, cfg.DatabaseURL)
	}

	expectedServerPort := ":8080"
	if cfg.ServerPort != expectedServerPort {
		t.Errorf("Expected ServerPort %s, got %s", expectedServerPort, cfg.ServerPort)
	}
}
