
CREATE TYPE status_order AS ENUM (
  'DRAFT',
  'CREATED'
);

CREATE TABLE IF NOT EXISTS orders (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    merchant_id uuid NOT NULL,
    FOREIGN KEY (merchant_id) REFERENCES merchants(id),
    detail jsonb NOT NULL,
    order_status status_order NOT NULL,
    location_lat FLOAT NOT NULL,
    location_long FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);