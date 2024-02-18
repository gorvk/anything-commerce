DROP TABLE IF EXISTS Products CASCADE;
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