DROP TABLE IF EXISTS Users CASCADE;
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(15) NOT NULL,
    last_name VARCHAR(15) NOT NULL,
    email VARCHAR(20) UNIQUE NOT NULL,
    phone_number VARCHAR(20) UNIQUE,
    user_address VARCHAR NOT NULL,
    is_seller BOOLEAN DEFAULT FALSE
);