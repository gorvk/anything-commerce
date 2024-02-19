CREATE OR REPLACE FUNCTION read_all_user() RETURNS TABLE (
        id INTEGER,
        _first_name VARCHAR,
        _last_name VARCHAR,
        _email VARCHAR,
        _phone_number VARCHAR,
        _user_address VARCHAR,
        is_seller BOOLEAN
    ) LANGUAGE plpgsql AS $$ BEGIN RETURN QUERY
SELECT *
FROM Users;
END;
$$;
CREATE OR REPLACE FUNCTION read_user_by_column(
        IN key_column_name TEXT,
        IN key_value ANYELEMENT
    ) RETURNS TABLE (
        id INTEGER,
        _first_name VARCHAR,
        _last_name VARCHAR,
        _email VARCHAR,
        _phone_number VARCHAR,
        _user_address VARCHAR,
        is_seller BOOLEAN
    ) LANGUAGE plpgsql AS $$ BEGIN RETURN QUERY EXECUTE format(
        'SELECT * FROM Users WHERE %I = $1',
        key_column_name
    ) USING key_value;
END;
$$;
-- SELECT * from my_function()