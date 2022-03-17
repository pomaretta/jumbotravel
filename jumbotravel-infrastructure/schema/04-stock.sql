CREATE DATABASE IF NOT EXISTS jumbotravel;
USE jumbotravel;

CREATE TABLE IF NOT EXISTS airplane_stock (

    -- PRIMARY KEY
    airplane_id         INT NOT NULL,
    product_id          INT NOT NULL,

    -- METADATA
    stock               INT NOT NULL,

    -- FLAGS (for future use)
    created_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    -- CONSTRAINT
    PRIMARY KEY (airplane_id, product_id),
    FOREIGN KEY (airplane_id) REFERENCES master_airplanes(airplane_id),
    FOREIGN KEY (product_id) REFERENCES master_products(product_id)
);