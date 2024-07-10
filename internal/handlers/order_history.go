package handlers

import (
	"encoding/json"
	"net/http"
	"statistics-collection/internal/database"
	"statistics-collection/internal/models"
)

// GetOrderHistoryHandler обрабатывает запрос на получение истории заказов для указанного клиента.
func GetOrderHistoryHandler(w http.ResponseWriter, r *http.Request) {
	clientName := r.URL.Query().Get("client_name")
	exchangeName := r.URL.Query().Get("exchange_name")
	label := r.URL.Query().Get("label")
	pair := r.URL.Query().Get("pair")

	client := &models.Client{
		ClientName:   clientName,
		ExchangeName: exchangeName,
		Label:        label,
		Pair:         pair,
	}

	orderHistory, err := database.GetOrderHistory(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderHistory)
}

// SaveOrderHandler обрабатывает запрос на сохранение заказа.
func SaveOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.HistoryOrder
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	if order.ClientName == "" || order.ExchangeName == "" || order.Label == "" || order.Pair == "" || order.Side == "" || order.Type == "" || order.BaseQty == 0 || order.Price == 0 {
		http.Error(w, "All fields must be provided", http.StatusBadRequest)
		return
	}

	if err := database.SaveOrderHistory(&order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
