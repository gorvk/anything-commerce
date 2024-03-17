CREATE OR REPLACE PROCEDURE update_order(
        IN column_name TEXT,
        IN key_column_name TEXT,
        IN new_value TEXT,
        IN key_value INTEGER
    ) LANGUAGE plpgsql AS $$ BEGIN 
    IF column_name = 'id' OR column_name = 'product_id' OR column_name = 'buyer_id' OR column_name = 'shop_id' THEN
        EXECUTE format('UPDATE Orders SET %I = $1::INTEGER WHERE %I = $2', column_name, key_column_name) USING new_value, key_value;
    ELSE
        EXECUTE format('UPDATE Orders SET %I = $1 WHERE %I = $2', column_name, key_column_name) USING new_value, key_value;
    END IF;
END; $$;