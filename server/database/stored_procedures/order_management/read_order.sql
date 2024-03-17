CREATE OR REPLACE FUNCTION read_all_orders() RETURNS TABLE (
        id INTEGER,
        from_map_location VARCHAR,
        to_map_location VARCHAR,
        last_stop_map_location VARCHAR,
        order_status VARCHAR,
        payment_status VARCHAR,
        product_id INTEGER,
        buyer_id INTEGER,
        shop_id INTEGER
    ) LANGUAGE plpgsql AS 
    $$ BEGIN 
        RETURN QUERY SELECT * FROM Orders;
    END; $$;

CREATE OR REPLACE FUNCTION read_order_by_column(
        IN key_column_name TEXT,
        IN key_value TEXT
    ) RETURNS TABLE (
        id INTEGER,
        from_map_location VARCHAR,
        to_map_location VARCHAR,
        last_stop_map_location VARCHAR,
        order_status VARCHAR,
        payment_status VARCHAR,
        product_id INTEGER,
        buyer_id INTEGER,
        shop_id INTEGER
    ) LANGUAGE plpgsql AS 
    $$ BEGIN 
        IF key_column_name = 'id' OR key_column_name = 'product_id' OR key_column_name = 'buyer_id' OR key_column_name = 'shop_id' THEN
            RETURN QUERY EXECUTE format('SELECT * FROM Orders WHERE %I = $1::INTEGER', key_column_name) USING key_value;
        ELSE
            RETURN QUERY EXECUTE format('SELECT * FROM Orders WHERE %I = $1', key_column_name) USING key_value;
        END IF;
    END; $$;