DROP TABLE IF EXISTS Shops CASCADE;
DROP TABLE IF EXISTS Users CASCADE;
DROP TABLE IF EXISTS Products CASCADE;
DROP TABLE IF EXISTS Orders CASCADE;
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(20) UNIQUE NOT NULL,
    phone_number varchar(20) UNIQUE,
    address varchar(50) NOT NULL,
    shop_name varchar(20),
    shop_id SERIAL,
    is_seller BOOLEAN DEFAULT FALSE
);
CREATE TABLE Shops (
    id SERIAL PRIMARY KEY,
    owner_id SERIAL NOT NULL,
    shop_name VARCHAR(20) NOT NULL,
    email VARCHAR(20) UNIQUE NOT NULL,
    phone_number varchar(20) UNIQUE,
    map_location varchar(50) NOT NULL,
    shop_type varchar(10) NOT NULL,
    shop_description varchar(200) NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES Users(id)
);
CREATE TABLE Products(
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(20) NOT NULL,
    seller_id SERIAL NOT NULL,
    shop_id SERIAL NOT NULL,
    product_type varchar(10),
    product_condition varchar(10),
    price MONEY,
    original_purchased_date DATE,
    original_purchaising_reciept_no varchar,
    product_description varchar(200),
    FOREIGN KEY (seller_id) REFERENCES Users(id),
    FOREIGN KEY (shop_id) REFERENCES Shops(id)
);
CREATE TABLE Orders(
    id SERIAL PRIMARY KEY,
    from_map_location VARCHAR NOT NULL,
    to_map_location VARCHAR NOT NULL,
    last_stop_map_location VARCHAR NOT NULL,
    order_status VARCHAR(10) NOT NULL,
    payment_status VARCHAR NOT NULL,
    product_id SERIAL NOT NULL,
    buyer_id SERIAL NOT NULL,
    shop_id SERIAL NOT NULL,
    FOREIGN KEY (product_id) REFERENCES Products(id),
    FOREIGN KEY (buyer_id) REFERENCES Users(id),
    FOREIGN KEY (shop_id) REFERENCES Shops(id)
);