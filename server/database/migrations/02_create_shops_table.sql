DROP TABLE IF EXISTS Shops CASCADE;
CREATE TABLE Shops (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER NOT NULL,
    shop_name VARCHAR(20) NOT NULL,
    email VARCHAR(20) UNIQUE NOT NULL,
    phone_number VARCHAR(20) UNIQUE,
    map_location VARCHAR(50) NOT NULL,
    shop_type VARCHAR(10) NOT NULL,
    shop_description VARCHAR(200) NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES Users(id)
);