CREATE OR REPLACE PROCEDURE delete_order(
        IN key_column_name TEXT,
        IN key_value TEXT
    ) LANGUAGE plpgsql AS $$ BEGIN 
    IF key_column_name = 'id' OR key_column_name = 'product_id' OR key_column_name = 'buyer_id' OR key_column_name = 'shop_id' THEN
        EXECUTE format('DELETE FROM Orders WHERE %I = $1::INTEGER', key_column_name) USING key_value;
    ELSE
        EXECUTE format('DELETE FROM Orders WHERE %I = $1', key_column_name) USING key_value;
    END IF;
END; $$;