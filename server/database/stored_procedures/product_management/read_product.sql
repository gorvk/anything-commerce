CREATE OR REPLACE FUNCTION read_all_products() RETURNS TABLE (
        _id INTEGER,
        _product_name VARCHAR,
        _seller_id INTEGER,
        _shop_id INTEGER,
        _product_type VARCHAR,
        _product_condition VARCHAR,
        _price MONEY,
        _original_purchased_date DATE,
        _original_purchaising_reciept_no VARCHAR,
        _product_description VARCHAR
    ) LANGUAGE plpgsql AS $$ BEGIN RETURN QUERY
SELECT *
FROM Products;
END;
$$;
CREATE OR REPLACE FUNCTION read_products_by_column(
        IN key_column_name TEXT,
        IN key_value ANYELEMENT
    ) RETURNS TABLE (
        _id INTEGER,
        _product_name VARCHAR,
        _seller_id INTEGER,
        _shop_id INTEGER,
        _product_type VARCHAR,
        _product_condition VARCHAR,
        _price MONEY,
        _original_purchased_date DATE,
        _original_purchaising_reciept_no VARCHAR,
        _product_description VARCHAR
    ) LANGUAGE plpgsql AS $$ BEGIN RETURN QUERY EXECUTE format(
        'SELECT * FROM Products WHERE %I = $1',
        key_column_name
    ) USING key_value;
END;
$$;
-- SELECT * from my_function()