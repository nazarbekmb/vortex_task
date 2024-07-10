package database

import (
	"log"
	"statistics-collection/internal/models"
)

// GetOrderBook возвращает Order Book для указанной биржи и валютной пары.
func GetOrderBook(exchangeName, pair string) ([]*models.DepthOrder, error) {
	var orders []*models.DepthOrder

	// Выполнение SQL запроса с JOIN
	query := `
        SELECT d.price, d.baseqty
        FROM depth_order d
        JOIN order_book o ON d.order_book_id = o.id
        WHERE o.exchange_name = $1 AND o.pair = $2;
    `
	rows, err := DB.Query(query, exchangeName, pair)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		return nil, err
	}
	defer rows.Close()

	// Обработка результатов запроса
	for rows.Next() {
		var order models.DepthOrder
		if err := rows.Scan(&order.Price, &order.BaseQty); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		orders = append(orders, &order)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		return nil, err
	}

	return orders, nil
}

func SaveOrderBook(exchangeName, pair string, asks, bids []models.DepthOrder) error {
	// Начинаем транзакцию
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Вставляем данные в order_book
	insertOrderBookStmt, err := tx.Prepare("INSERT INTO order_book (exchange_name, pair) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return err
	}
	defer insertOrderBookStmt.Close()

	var orderBookID int
	err = insertOrderBookStmt.QueryRow(exchangeName, pair).Scan(&orderBookID)
	if err != nil {
		return err
	}

	// Вставляем данные в depth_order для asks
	insertDepthOrderStmt, err := tx.Prepare("INSERT INTO depth_order (order_book_id, side, price, baseqty) VALUES ($1, 'ask', $2, $3)")
	if err != nil {
		return err
	}
	defer insertDepthOrderStmt.Close()

	for _, order := range asks {
		_, err := insertDepthOrderStmt.Exec(orderBookID, order.Price, order.BaseQty)
		if err != nil {
			return err
		}
	}

	// Вставляем данные в depth_order для bids
	insertDepthOrderStmt, err = tx.Prepare("INSERT INTO depth_order (order_book_id, side, price, baseqty) VALUES ($1, 'bid', $2, $3)")
	if err != nil {
		return err
	}

	for _, order := range bids {
		_, err := insertDepthOrderStmt.Exec(orderBookID, order.Price, order.BaseQty)
		if err != nil {
			return err
		}
	}

	return nil
}
