CREATE OR REPLACE PROCEDURE delete_shop(
        IN key_column_name TEXT,
        IN key_value TEXT
    ) LANGUAGE plpgsql AS 
$$ BEGIN 
    IF key_column_name = 'id' OR key_column_name = 'owner_id' THEN
        EXECUTE format('DELETE FROM Shops WHERE %I = $1::INTEGER', key_column_name) USING key_value;
    ELSE 
        EXECUTE format('DELETE FROM Shops WHERE %I = $1', key_column_name) USING key_value;
    END IF;
END; $$;