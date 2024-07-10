package database

import (
	"database/sql"
	"log"
	"statistics-collection/internal/models"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

const testConnStr = "postgres://postgres:123456@localhost:5432/testdb?sslmode=disable"

func SetupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testConnStr)
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS order_book (
            id SERIAL PRIMARY KEY,
            exchange_name VARCHAR(50) NOT NULL,
            pair VARCHAR(50) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
        CREATE TABLE IF NOT EXISTS depth_order (
            id SERIAL PRIMARY KEY,
            order_book_id INT NOT NULL REFERENCES order_book(id) ON DELETE CASCADE,
            side VARCHAR(4) NOT NULL CHECK (side IN ('ask', 'bid')),
            price DOUBLE PRECISION NOT NULL,
            baseqty DOUBLE PRECISION NOT NULL
        );
        CREATE TABLE IF NOT EXISTS order_history (
            id SERIAL PRIMARY KEY,
            client_name VARCHAR(50) NOT NULL,
            exchange_name VARCHAR(50) NOT NULL,
            label VARCHAR(50) NOT NULL,
            pair VARCHAR(50) NOT NULL,
            side VARCHAR(4) NOT NULL,
            type VARCHAR(10) NOT NULL,
            base_qty DOUBLE PRECISION NOT NULL,
            price DOUBLE PRECISION NOT NULL,
            algorithm_name_placed VARCHAR(50),
            lowest_sell_prc DOUBLE PRECISION,
            highest_buy_prc DOUBLE PRECISION,
            commission_quote_qty DOUBLE PRECISION,
            time_placed TIMESTAMP NOT NULL
        );
        INSERT INTO order_book (exchange_name, pair) VALUES ('Binance', 'BTC/USDT');
        INSERT INTO depth_order (order_book_id, side, price, baseqty) VALUES (1, 'ask', 45000.0, 0.5), (1, 'bid', 44000.0, 1.0);
    `)
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}

	return db
}

func TeardownTestDB(t *testing.T, db *sql.DB) {
	_, err := db.Exec(`
        DROP TABLE IF EXISTS depth_order;
        DROP TABLE IF EXISTS order_book;
        DROP TABLE IF EXISTS order_history;
    `)
	if err != nil {
		t.Fatalf("Failed to tear down test database: %v", err)
	}
	db.Close()
}

func TestGetOrderBook(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	DB = db // Используем тестовую базу данных

	orders, err := GetOrderBook("Binance", "BTC/USDT")
	if err != nil {
		t.Fatalf("Failed to get order book: %v", err)
	}
	log.Println(orders)
	if len(orders) != 2 {
		t.Fatalf("Expected 2 orders, got %d", len(orders))
	}

	expectedOrders := []*models.DepthOrder{
		{Price: 45000.0, BaseQty: 0.5},
		{Price: 44000.0, BaseQty: 1.0},
	}

	for i, order := range orders {
		if order.Price != expectedOrders[i].Price || order.BaseQty != expectedOrders[i].BaseQty {
			t.Errorf("Expected order %+v, got %+v", expectedOrders[i], order)
		}
	}
}

func TestSaveOrderHistory(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	DB = db // Используем тестовую базу данных

	timePlaced, err := time.Parse("2006-01-02 15:04:05", "2023-07-10 12:00:00")
	if err != nil {
		t.Fatalf("Failed to parse time: %v", err)
	}

	order := &models.HistoryOrder{
		ClientName:          "test_client",
		ExchangeName:        "Binance",
		Label:               "test_label",
		Pair:                "BTC/USDT",
		Side:                "buy",
		Type:                "limit",
		BaseQty:             1.0,
		Price:               45000.0,
		AlgorithmNamePlaced: "test_algo",
		LowestSellPrice:     45050.0,
		HighestBuyPrice:     44950.0,
		CommissionQuoteQty:  0.001,
		TimePlaced:          timePlaced,
	}

	err = SaveOrderHistory(order)
	if err != nil {
		t.Fatalf("Failed to save order history: %v", err)
	}

	var count int
	err = DB.QueryRow(`SELECT COUNT(*) FROM order_history WHERE client_name = $1 AND exchange_name = $2 AND label = $3 AND pair = $4`,
		order.ClientName, order.ExchangeName, order.Label, order.Pair).Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query order history: %v", err)
	}

	if count != 1 {
		t.Fatalf("Expected 1 order history record, got %d", count)
	}
}
