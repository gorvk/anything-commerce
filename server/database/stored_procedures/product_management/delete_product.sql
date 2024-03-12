CREATE OR REPLACE PROCEDURE delete_product(
        IN key_column_name TEXT,
        IN key_value ANYELEMENT
    ) LANGUAGE plpgsql AS $$ BEGIN EXECUTE format(
        'DELETE FROM Products WHERE %I = $1',
        key_column_name
    ) USING key_value;
END;
$$;