CREATE OR REPLACE PROCEDURE delete_user(
        IN key_column_name TEXT,
        IN key_value TEXT
    ) LANGUAGE plpgsql AS 
$$ BEGIN
    IF key_column_name = 'id' THEN
        EXECUTE format('DELETE FROM Users WHERE %I = $1::INTEGER', key_column_name) USING key_value;
    ELSEIF key_column_name = 'is_seller' THEN
        EXECUTE format('DELETE FROM Users WHERE %I = $1::BOOLEAN', key_column_name) USING key_value;
    ELSE 
        EXECUTE format('DELETE FROM Users WHERE %I = $1', key_column_name) USING key_value;
    END IF;
END; $$;