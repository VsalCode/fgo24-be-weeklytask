-- Active: 1750752675497@@127.0.0.1@5432@ewalletdb
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    pin INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE wallets (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    balance DECIMAL(15, 2) DEFAULT 0.00,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE payment_method (
    id SERIAL PRIMARY KEY,
    method_name VARCHAR(100) NOT NULL
);

INSERT INTO payment_method (method_name) VALUES
('bri'),
('dana'),
('bca'),
('gopay'),
('ovo');

CREATE TABLE topup (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    topup_amount DECIMAL(15, 2) NOT NULL,
    topup_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    method_id INT REFERENCES payment_method(id),
    success BOOLEAN DEFAULT FALSE
);

CREATE TABLE transfers (
    transfer_id SERIAL PRIMARY KEY,
    sender_user_id INT REFERENCES users(id) ON DELETE CASCADE,
    receiver_user_id INT REFERENCES users(id) ON DELETE CASCADE,
    transfer_amount DECIMAL(15, 2) NOT NULL,
    transfer_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    success BOOLEAN DEFAULT FALSE
);
