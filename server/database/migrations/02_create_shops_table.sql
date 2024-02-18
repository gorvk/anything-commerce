DROP TABLE IF EXISTS Shops CASCADE;
CREATE TABLE Shops (
    id SERIAL PRIMARY KEY,
    owner_id SERIAL NOT NULL,
    shop_name VARCHAR(20) NOT NULL,
    email VARCHAR(20) UNIQUE NOT NULL,
    phone_number varchar(20) UNIQUE,
    map_location varchar(50) NOT NULL,
    shop_type varchar(10) NOT NULL,
    shop_description varchar(200) NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES Users(id)
);