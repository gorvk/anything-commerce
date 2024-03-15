CREATE OR REPLACE PROCEDURE update_user(
	IN t_column_name text,
	IN key_column_name text,
	IN new_value text,
	IN key_value text)
LANGUAGE 'plpgsql'
AS $BODY$
 
BEGIN
    IF t_column_name = 'is_seller' THEN
        BEGIN
            EXECUTE format(
                'UPDATE Users SET %I = $1::bool WHERE %I = $2',
                t_column_name,
                key_column_name
            ) USING new_value, key_value;
        EXCEPTION
            WHEN others THEN
                RAISE NOTICE 'Could not update is_seller with value: %', new_value;
        END;
    ELSE 
        EXECUTE format(
            'UPDATE Users SET %I = $1 WHERE %I = $2',
            t_column_name,
            key_column_name
        ) USING new_value, key_value;
    END IF;
END;
$BODY$;