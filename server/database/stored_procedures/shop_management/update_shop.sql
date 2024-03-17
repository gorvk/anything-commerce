CREATE OR REPLACE PROCEDURE update_shop(
    IN column_name TEXT,
    IN key_column_name TEXT,
    IN new_value TEXT,
    IN key_value TEXT
) LANGUAGE plpgsql AS 
$$ BEGIN 
    IF column_name = 'id' OR column_name = 'owner_id' THEN
        EXECUTE format('UPDATE Shops SET %I = $1::INTEGER WHERE %I = $2', column_name, key_column_name) USING new_value, key_value;
    ELSE 
        EXECUTE format('UPDATE Shops SET %I = $1 WHERE %I = $2', column_name, key_column_name) USING new_value, key_value;
    END IF;
END; $$;