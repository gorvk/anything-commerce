DROP TABLE IF EXISTS Users CASCADE;
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(20) UNIQUE NOT NULL,
    phone_number varchar(20) UNIQUE,
    address varchar(50) NOT NULL,
    shop_name varchar(20),
    shop_id SERIAL,
    is_seller BOOLEAN DEFAULT FALSE
);