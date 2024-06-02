
CREATE TYPE status_order AS ENUM (
  'DRAFT',
  'CREATED',
  'PURCHASED'
);

CREATE TABLE IF NOT EXISTS orders (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    order_status status_order NOT NULL,
    location_lat FLOAT NOT NULL,
    location_long FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);