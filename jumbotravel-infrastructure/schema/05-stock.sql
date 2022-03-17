CREATE DATABASE IF NOT EXISTS jumbotravel;
USE jumbotravel;

CREATE TABLE IF NOT EXISTS airplane_stock (

    -- PRIMARY KEY
    stock_id   INT NOT NULL AUTO_INCREMENT,

    airplane_id         INT NOT NULL,
    airplanemapping_id  INT NOT NULL,
    product_id          INT NOT NULL,
    productmapping_id   INT NOT NULL,

    -- METADATA
    stock               INT NOT NULL,

    -- FLAGS (for future use)
    created_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    -- CONSTRAINT
    PRIMARY KEY (stock_id),
    FOREIGN KEY (airplane_id) REFERENCES master_airplanes(airplane_id),
    FOREIGN KEY (airplanemapping_id) REFERENCES master_airplanesmapping(airplanemapping_id),
    FOREIGN KEY (product_id) REFERENCES master_products(product_id),
    FOREIGN KEY (productmapping_id) REFERENCES master_productmapping(productmapping_id)
);

