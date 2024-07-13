-- Active: 1718919020656@@127.0.0.1@5432@master
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    user_type VARCHAR(20) NOT NULL,
    address TEXT,
    phone_number VARCHAR(20),
    bio TEXT,
    specialties TEXT,
    years_of_experience INTEGER,
    is_verified BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

create table refresh_token(
    id uuid primary key DEFAULT gen_random_uuid() not null,
    email text not null,
    user_id uuid references users(id),
    token text UNIQUE NOT NULL,
    expires_at bigint,
    created_at BIGINT,
    deleted_at BIGINT
);