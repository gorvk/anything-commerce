CREATE OR REPLACE FUNCTION read_all_shops() RETURNS TABLE (
        id INTEGER,
        owner_id INTEGER,
        shop_name VARCHAR,
        email VARCHAR,
        phone_number VARCHAR,
        map_location VARCHAR,
        shop_type VARCHAR,
        shop_description VARCHAR
    ) LANGUAGE plpgsql AS 
    $$ BEGIN 
        RETURN QUERY SELECT * FROM Shops;
    END $$;

CREATE OR REPLACE FUNCTION read_shop_by_column(
        IN key_column_name TEXT,
        IN key_value TEXT
    ) RETURNS TABLE (
        id INTEGER,
			owner_id INTEGER,
        shop_name VARCHAR,
        email VARCHAR,
        phone_number VARCHAR,
        map_location VARCHAR,
        shop_type VARCHAR,
        shop_description VARCHAR
    ) LANGUAGE plpgsql 
AS $$ 
    BEGIN 
        IF key_column_name = 'id'OR key_column_name = 'owner_id' THEN
            RETURN QUERY EXECUTE format('SELECT * FROM Shops WHERE %I = $1::INTEGER', key_column_name) USING key_value;
        ELSE 
            RETURN QUERY EXECUTE format('SELECT * FROM Shops WHERE %I = $1', key_column_name) USING key_value;
        END IF;
END; $$;