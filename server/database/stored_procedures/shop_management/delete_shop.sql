CREATE OR REPLACE PROCEDURE delete_shop(
        IN key_column_name TEXT,
        IN key_value ANYELEMENT
    ) LANGUAGE plpgsql AS $$ BEGIN EXECUTE format(
        'DELETE FROM Shops WHERE %I = $1',
        key_column_name
    ) USING key_value;
END;
$$;