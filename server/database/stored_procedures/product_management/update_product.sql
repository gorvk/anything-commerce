CREATE OR REPLACE PROCEDURE update_product(
        IN column_name TEXT,
        IN key_column_name TEXT,
        IN new_value TEXT,
        IN key_value INTEGER
    ) LANGUAGE plpgsql AS 
    $$ BEGIN 
        IF column_name = 'id' OR column_name = 'seller_id' OR column_name = 'shop_id' THEN
            EXECUTE format('UPDATE Products SET %I = $1::INTEGER WHERE %I = $2', column_name, key_column_name) USING new_value, key_value;
        ELSEIF column_name = 'price' THEN
            EXECUTE format('UPDATE Products SET %I = $1::MONEY WHERE %I = $2', column_name, key_column_name) USING new_value, key_value;
        ELSEIF column_name = 'original_purchased_date' THEN
            EXECUTE format('UPDATE Products SET %I = $1::DATE WHERE %I = $2', column_name, key_column_name) USING new_value, key_value;
        ELSE 
            EXECUTE format('UPDATE Products SET %I = $1 WHERE %I = $2', column_name, key_column_name) USING new_value, key_value;    
        END IF;
    END; $$;