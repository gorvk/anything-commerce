CREATE OR REPLACE FUNCTION read_all_products() RETURNS TABLE (
        id INTEGER,
        product_name VARCHAR,
        seller_id INTEGER,
        shop_id INTEGER,
        product_type VARCHAR,
        product_condition VARCHAR,
        price MONEY,
        original_purchased_date DATE,
        original_purchaising_reciept_no VARCHAR,
        product_description VARCHAR
    ) LANGUAGE plpgsql AS 
    $$ BEGIN RETURN QUERY
        SELECT * FROM Products;
    END; $$;


CREATE OR REPLACE FUNCTION read_product_by_column(
        IN key_column_name TEXT,
        IN key_value TEXT
    ) RETURNS TABLE (
        id INTEGER,
        product_name VARCHAR,
        seller_id INTEGER,
        shop_id INTEGER,
        product_type VARCHAR,
        product_condition VARCHAR,
        price MONEY,
        original_purchased_date DATE,
        original_purchaising_reciept_no VARCHAR,
        product_description VARCHAR
    ) LANGUAGE plpgsql AS 
    $$ BEGIN 
        IF key_column_name = 'id' OR key_column_name = 'seller_id' OR key_column_name = 'shop_id' THEN
            RETURN QUERY EXECUTE format('SELECT * FROM Products WHERE %I = $1::INTEGER', key_column_name) USING key_value;
        ELSEIF key_column_name = 'price' THEN
            RETURN QUERY EXECUTE format('SELECT * FROM Products WHERE %I = $1::MONEY', key_column_name) USING key_value;
        ELSEIF key_column_name = 'original_purchased_date' THEN
            RETURN QUERY EXECUTE format('SELECT * FROM Products WHERE %I = $1::DATE', key_column_name) USING key_value;
        ELSE 
            RETURN QUERY EXECUTE format('SELECT * FROM Products WHERE %I = $1', key_column_name) USING key_value;
        END IF;
    END; $$;