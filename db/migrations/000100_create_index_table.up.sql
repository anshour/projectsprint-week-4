-- CREATE INDEX idx_users_nip ON users (nip);

CREATE INDEX idx_merchants_on_created_at ON merchants(created_at);
CREATE INDEX idx_merchant_items_on_created_at ON merchant_items(created_at);

