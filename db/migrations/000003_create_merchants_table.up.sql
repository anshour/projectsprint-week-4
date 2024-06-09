CREATE TYPE merchant_category AS ENUM (
    'SmallRestaurant',
	'MediumRestaurant',
	'LargeRestaurant',
	'MerchandiseRestaurant',
	'BoothKiosk',
	'ConvenienceStore'
);

CREATE TABLE IF NOT EXISTS merchants (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(30) NOT NULL,
    category merchant_category NOT NULL,
    image_url VARCHAR(225) NOT NULL,
    location_lat FLOAT NOT NULL,
    location_long FLOAT NOT NULL,
    geo_hash VARCHAR(12) not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);