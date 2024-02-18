DROP TABLE IF EXISTS Orders CASCADE;
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