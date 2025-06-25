-- E wallet

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone VARCHAR(30) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    pin INT NOT NULL CHECK (pin BETWEEN 1000 AND 9999),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    token VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expired_at TIMESTAMP,
    user_id INT REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE payment_methods (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULLj
);

CREATE TYPE transaction_type AS ENUM ('topup', 'transfer');

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    amount DECIMAL(15, 2) NOT NULL CHECK (amount > 0),
    type transaction_type,
    payment_method_id INT REFERENCES payment_methods(id) ON DELETE SET NULL,
    user_id INT REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE wallets (
    id SERIAL PRIMARY KEY,
    balance DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
    user_id INT REFERENCES users(id) ON DELETE CASCADE
);
