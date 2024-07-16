-- Active: 1718919020656@@127.0.0.1@5432@restaurants@public

CREATE TABLE kitchens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    cuisine_type VARCHAR(50),
    address TEXT NOT NULL,
    phone_number VARCHAR(20),
    rating DECIMAL(3, 2) DEFAULT 0,
    total_orders INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP 
);

CREATE TABLE dishes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    kitchen_id UUID REFERENCES kitchens(id),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    category VARCHAR(50),
    ingredients TEXT[],
    allergens TEXT[],
    nutrition_info JSONB,
    dietary_info TEXT[],
    available BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    kitchen_id UUID REFERENCES kitchens(id),
    items JSONB NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',  -- Example default value 'pending'
    delivery_address TEXT NOT NULL,
    delivery_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID REFERENCES orders(id),
    user_id UUID NOT NULL,
    kitchen_id UUID REFERENCES kitchens(id),
    rating DECIMAL(2, 1) NOT NULL,
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ,
    deleted_at TIMESTAMP
);

CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID REFERENCES orders(id),
    amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    payment_method VARCHAR(50) NOT NULL,
    transaction_id VARCHAR(100),
    card_number VARCHAR(20),
    expiry_date VARCHAR(7),
    cvv VARCHAR(4),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE working_hours (
    kitchen_id UUID REFERENCES kitchens(id),
    day_of_week VARCHAR(20) NOT NULL,
    open_time TIME NOT NULL,
    close_time TIME NOT NULL,
    PRIMARY KEY (kitchen_id, day_of_week)
);

CREATE TABLE user_preferences (
    user_id UUID NOT NULL,
    cuisine_type VARCHAR(50),
    dietary_preferences TEXT[],
    favorite_kitchen_ids UUID[],
    PRIMARY KEY (user_id)
);
