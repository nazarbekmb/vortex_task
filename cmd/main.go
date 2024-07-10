package main

import (
	"log"
	"net/http"
	"statistics-collection/internal/config"
	"statistics-collection/internal/database"
	"statistics-collection/internal/migrations"
	"statistics-collection/internal/routers"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Загрузка конфигураций
	cfg := config.LoadConfig()

	// Инициализация базы данных
	if err := database.InitDB(cfg.DatabaseURL); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer database.DB.Close()

	// Применение миграций
	if err := migrations.ApplyMigrations(database.DB); err != nil {
		log.Fatalf("Could not apply migrations: %v", err)
	}

	// Регистрация обработчиков HTTP запросов
	routers.InitRoutes()

	// Запуск сервера на порту 8080
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
