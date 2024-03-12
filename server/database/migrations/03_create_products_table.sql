DROP TABLE IF EXISTS Products CASCADE;
CREATE TABLE Products(
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(20) NOT NULL,
    seller_id SERIAL NOT NULL,
    shop_id SERIAL NOT NULL,
    product_type VARCHAR(10),
    product_condition VARCHAR(10),
    price MONEY,
    original_purchased_date DATE,
    original_purchaising_reciept_no VARCHAR,
    product_description VARCHAR(200),
    FOREIGN KEY (seller_id) REFERENCES Users(id),
    FOREIGN KEY (shop_id) REFERENCES Shops(id)
);