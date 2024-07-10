package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// DB хранит соединение с базой данных PostgreSQL.
var DB *sql.DB

// InitDB инициализирует соединение с PostgreSQL.
func InitDB(connStr string) error {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	log.Println("Connected to PostgreSQL database")
	return nil
}
