CREATE OR REPLACE FUNCTION read_all_orders() RETURNS TABLE (
        _id INTEGER,
        _from_map_location VARCHAR,
        _to_map_location VARCHAR,
        _last_stop_map_location VARCHAR,
        _order_status VARCHAR,
        _payment_status VARCHAR,
        _product_id INTEGER,
        _buyer_id INTEGER,
        _shop_id INTEGER
    ) LANGUAGE plpgsql AS $$ BEGIN RETURN QUERY
SELECT *
FROM Orders;
END;
$$;
CREATE OR REPLACE FUNCTION read_order_by_column(
        IN key_column_name TEXT,
        IN key_value ANYELEMENT
    ) RETURNS TABLE (
        _id INTEGER,
        _from_map_location VARCHAR,
        _to_map_location VARCHAR,
        _last_stop_map_location VARCHAR,
        _order_status VARCHAR,
        _payment_status VARCHAR,
        _product_id INTEGER,
        _buyer_id INTEGER,
        _shop_id INTEGER
    ) LANGUAGE plpgsql AS $$ BEGIN RETURN QUERY EXECUTE format(
        'SELECT * FROM Orders WHERE %I = $1',
        key_column_name
    ) USING key_value;
END;
$$;
-- SELECT * from my_function()