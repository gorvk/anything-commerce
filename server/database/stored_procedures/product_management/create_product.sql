CREATE OR REPLACE PROCEDURE create_product (
        _product_name VARCHAR,
        _seller_id INTEGER,
        _shop_id INTEGER,
        _product_type VARCHAR,
        _product_condition VARCHAR,
        _price MONEY,
        _original_purchased_date DATE,
        _original_purchaising_reciept_no VARCHAR,
        _product_description VARCHAR
    ) LANGUAGE SQL AS $$
INSERT INTO Products (
        product_name,
        seller_id,
        shop_id,
        product_type,
        product_condition,
        price,
        original_purchased_date,
        original_purchaising_reciept_no,
        product_description
    )
VALUES (
        _product_name,
        _seller_id,
        _shop_id,
        _product_type,
        _product_condition,
        _price,
        _original_purchased_date,
        _original_purchaising_reciept_no,
        _product_description
    );
$$