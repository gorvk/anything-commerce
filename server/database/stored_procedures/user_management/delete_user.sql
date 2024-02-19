CREATE OR REPLACE PROCEDURE delete_user(
        IN key_column_name TEXT,
        IN key_value ANYELEMENT
    ) LANGUAGE plpgsql AS $$ BEGIN EXECUTE format(
        'DELETE FROM Users WHERE %I = $1',
        key_column_name
    ) USING key_value;
END;
$$;