
CREATE TABLE IF NOT EXISTS merhcant_item_orders (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    merchant_order_id uuid NOT NULL,
    FOREIGN KEY (merchant_order_id) REFERENCES merchant_orders(id),
    quantity INTEGER NOT NULL,
    merchant_item_id uuid NOT NULL,
    FOREIGN KEY (merchant_item_id) REFERENCES merchant_items(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);