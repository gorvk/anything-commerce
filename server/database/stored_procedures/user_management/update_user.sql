CREATE OR REPLACE PROCEDURE update_user(
	IN column_name text,
	IN key_column_name text,
	IN new_value text,
	IN key_value text)
LANGUAGE 'plpgsql' AS 
$$ BEGIN
    IF column_name = 'id' THEN
        EXECUTE format('UPDATE Users SET %I = $1::INTEGER WHERE %I = $2', column_name, key_column_name) USING new_value, key_value;
    ELSEIF column_name = 'is_seller' THEN
        EXECUTE format('UPDATE Users SET %I = $1::BOOLEAN WHERE %I = $2', column_name, key_column_name) USING new_value, key_value;
    ELSE 
        EXECUTE format('UPDATE Users SET %I = $1 WHERE %I = $2', column_name, key_column_name) USING new_value, key_value;
    END IF;
END; $$;