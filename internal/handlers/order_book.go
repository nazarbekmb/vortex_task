package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"statistics-collection/internal/database"
	"statistics-collection/internal/models"
)

// GetOrderBookHandler обрабатывает запрос на получение Order Book.
func GetOrderBookHandler(w http.ResponseWriter, r *http.Request) {
	exchangeName := r.URL.Query().Get("exchange_name")
	pair := r.URL.Query().Get("pair")

	if exchangeName == "" || pair == "" {
		http.Error(w, "Exchange name and pair must be provided", http.StatusBadRequest)
		return
	}

	orderBook, err := database.GetOrderBook(exchangeName, pair)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderBook)
}

// SaveOrderBookHandler обрабатывает запрос на сохранение Order Book.
func SaveOrderBookHandler(w http.ResponseWriter, r *http.Request) {
	var orderBook models.OrderBook
	if err := json.NewDecoder(r.Body).Decode(&orderBook); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Преобразование []*models.DepthOrder в []models.DepthOrder
	asks := make([]models.DepthOrder, len(orderBook.Asks))
	for i, ask := range orderBook.Asks {
		asks[i] = *ask
	}

	bids := make([]models.DepthOrder, len(orderBook.Bids))
	for i, bid := range orderBook.Bids {
		bids[i] = *bid
	}

	if orderBook.Exchange == "" || orderBook.Pair == "" {
		http.Error(w, "Exchange name and pair must be provided", http.StatusBadRequest)
		return
	}

	if err := database.SaveOrderBook(orderBook.Exchange, orderBook.Pair, asks, bids); err != nil {
		log.Printf("Error saving order book: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
