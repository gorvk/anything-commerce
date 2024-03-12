CREATE OR REPLACE FUNCTION read_all_shops() RETURNS TABLE (
        _id INTEGER,
        _shop_name VARCHAR,
        _email VARCHAR,
        _phone_number VARCHAR,
        _map_location VARCHAR,
        _shop_type VARCHAR,
        _shop_description VARCHAR
    ) LANGUAGE plpgsql AS $$ BEGIN RETURN QUERY
SELECT *
FROM Shops;
END;
$$;
CREATE OR REPLACE FUNCTION read_shop_by_column(
        IN key_column_name TEXT,
        IN key_value ANYELEMENT
    ) RETURNS TABLE (
        _id INTEGER,
        _shop_name VARCHAR,
        _email VARCHAR,
        _phone_number VARCHAR,
        _map_location VARCHAR,
        _shop_type VARCHAR,
        _shop_description VARCHAR
    ) LANGUAGE plpgsql AS $$ BEGIN RETURN QUERY EXECUTE format(
        'SELECT * FROM Shops WHERE %I = $1',
        key_column_name
    ) USING key_value;
END;
$$;
-- SELECT * from my_function()