CREATE OR REPLACE PROCEDURE create_order (
        _from_map_location VARCHAR,
        _to_map_location VARCHAR,
        _last_stop_map_location VARCHAR,
        _order_status VARCHAR,
        _payment_status VARCHAR,
        _product_id INTEGER,
        _buyer_id INTEGER,
        _shop_id INTEGER
    ) LANGUAGE SQL AS $$
INSERT INTO Orders (
        from_map_location,
        to_map_location,
        last_stop_map_location,
        order_status,
        payment_status,
        product_id,
        buyer_id,
        shop_id
    )
VALUES (
        _from_map_location,
        _to_map_location,
        _last_stop_map_location,
        _order_status,
        _payment_status,
        _product_id,
        _buyer_id,
        _shop_id
    );
$$