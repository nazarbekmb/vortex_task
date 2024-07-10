package database

import "statistics-collection/internal/models"

// GetOrderHistory возвращает историю заказов для указанного клиента.
func GetOrderHistory(client *models.Client) ([]*models.HistoryOrder, error) {
	var orders []*models.HistoryOrder

	query := `
		SELECT client_name, exchange_name, label, pair, side, type, base_qty, price,
		       algorithm_name_placed, lowest_sell_prc, highest_buy_prc, commission_quote_qty, time_placed
		FROM order_history
		WHERE client_name = $1 AND exchange_name = $2 AND label = $3 AND pair = $4
	`

	rows, err := DB.Query(query, client.ClientName, client.ExchangeName, client.Label, client.Pair)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.HistoryOrder
		if err := rows.Scan(
			&order.ClientName, &order.ExchangeName, &order.Label, &order.Pair, &order.Side, &order.Type,
			&order.BaseQty, &order.Price, &order.AlgorithmNamePlaced, &order.LowestSellPrice, &order.HighestBuyPrice,
			&order.CommissionQuoteQty, &order.TimePlaced,
		); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}

// SaveOrder сохраняет заказ для указанного клиента в базе данных.
func SaveOrderHistory(order *models.HistoryOrder) error {
	query := `
		INSERT INTO order_history (client_name, exchange_name, label, pair, side, type, base_qty, price,
		                           algorithm_name_placed, lowest_sell_prc, highest_buy_prc, commission_quote_qty, time_placed)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	_, err := DB.Exec(query, order.ClientName, order.ExchangeName, order.Label, order.Pair, order.Side, order.Type,
		order.BaseQty, order.Price, order.AlgorithmNamePlaced, order.LowestSellPrice, order.HighestBuyPrice,
		order.CommissionQuoteQty, order.TimePlaced)
	if err != nil {
		return err
	}

	return nil
}
