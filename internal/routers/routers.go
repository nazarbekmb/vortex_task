package routers

import (
	"net/http"
	"statistics-collection/internal/handlers"
)

// InitRoutes инициализирует маршруты для приложения
func InitRoutes() {
	http.HandleFunc("/api/get_order_book", handlers.GetOrderBookHandler)
	http.HandleFunc("/api/save_order_book", handlers.SaveOrderBookHandler)
	http.HandleFunc("/api/get_order_history", handlers.GetOrderHistoryHandler)
	http.HandleFunc("/api/save_order_history", handlers.SaveOrderHandler)
}
