CREATE OR REPLACE PROCEDURE edit_user(
        IN column_nam TEXT,
        IN key_column_name TEXT,
        IN new_value ANYELEMENT,
        IN key_value ANYELEMENT
    ) LANGUAGE plpgsql AS $$ BEGIN EXECUTE format(
        'UPDATE Users SET %I = $1 WHERE %I = $2',
        column_nam,
        key_column_name
    ) USING new_value,
    key_value;
END;
$$;