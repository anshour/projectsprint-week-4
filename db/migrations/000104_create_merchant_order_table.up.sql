CREATE TABLE IF NOT EXISTS merchant_orders (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id uuid NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    merchant_id uuid NOT NULL,
    FOREIGN KEY (merchant_id) REFERENCES merchants(id),
    item_id uuid NOT NULL,
    FOREIGN KEY (item_id) REFERENCES merchant_items(id),
    item_price INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);