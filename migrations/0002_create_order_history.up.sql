CREATE TABLE order_history (
    client_name              VARCHAR(100) NOT NULL,
    exchange_name            VARCHAR(50) NOT NULL,
    label                    VARCHAR(100),
    pair                     VARCHAR(50) NOT NULL,
    side                     VARCHAR(10) NOT NULL,
    type                     VARCHAR(20) NOT NULL,
    base_qty                 DOUBLE PRECISION NOT NULL,
    price                    DOUBLE PRECISION NOT NULL,
    algorithm_name_placed    VARCHAR(100),
    lowest_sell_prc          DOUBLE PRECISION,
    highest_buy_prc          DOUBLE PRECISION,
    commission_quote_qty     DOUBLE PRECISION,
    time_placed              TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
