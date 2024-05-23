CREATE TYPE merchant_item_category AS ENUM (
    'Beverage',
    'Food',
	'Snack',
	'Condiments',
	'Additions'
);

CREATE TABLE IF NOT EXISTS merchant_items(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(30) NOT NULL,
    merchant_id uuid NOT NULL,
    FOREIGN KEY (merchant_id) REFERENCES merchants(id),
    category merchant_item_category NOT NULL,
    price INTEGER NOT NULL,
    image_url VARCHAR(255) NOT NULL
);