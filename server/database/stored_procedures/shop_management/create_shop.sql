CREATE OR REPLACE PROCEDURE create_shop (
        _owner_id INTEGER,
        _shop_name VARCHAR,
        _email VARCHAR,
        _phone_number VARCHAR,
        _map_location VARCHAR,
        _shop_type VARCHAR,
        _shop_description VARCHAR
    ) LANGUAGE SQL AS $$
INSERT INTO Shops (
        owner_id,
        shop_name,
        email,
        phone_number,
        map_location,
        shop_type,
        shop_description
    )
VALUES (
        _owner_id,
        _shop_name,
        _email,
        _phone_number,
        _map_location,
        _shop_type,
        _shop_description
    );
$$