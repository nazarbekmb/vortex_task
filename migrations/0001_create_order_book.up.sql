CREATE TABLE order_book (
    id SERIAL PRIMARY KEY,
    exchange_name VARCHAR(50) NOT NULL,
    pair VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE depth_order (
    id SERIAL PRIMARY KEY,
    order_book_id INT NOT NULL REFERENCES order_book(id) ON DELETE CASCADE,
    side VARCHAR(4) NOT NULL CHECK (side IN ('ask', 'bid')), 
    price DOUBLE PRECISION NOT NULL,
    baseqty DOUBLE PRECISION NOT NULL
);
