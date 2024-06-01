CREATE TABLE IF NOT EXISTS estimations (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id uuid NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    total_price INT NOT NULL,
    estimation_minutes INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);