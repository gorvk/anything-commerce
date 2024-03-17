CREATE OR REPLACE PROCEDURE delete_product(
        IN key_column_name TEXT,
        IN key_value TEXT
    ) LANGUAGE plpgsql AS 
$$ BEGIN 
    IF key_column_name = 'id' OR key_column_name = 'seller_id' OR key_column_name = 'shop_id' THEN
        EXECUTE format('DELETE FROM Products WHERE %I = $1::INTEGER', key_column_name) USING key_value;
    ELSEIF key_column_name = 'price' THEN
        EXECUTE format('DELETE FROM Products WHERE %I = $1::MONEY', key_column_name) USING key_value;
    ELSEIF key_column_name = 'original_purchased_date' THEN
        EXECUTE format('DELETE FROM Products WHERE %I = $1::DATE', key_column_name) USING key_value;
    ELSE 
        EXECUTE format('DELETE FROM Products WHERE %I = $1', key_column_name) USING key_value;
    END IF;
END; $$;