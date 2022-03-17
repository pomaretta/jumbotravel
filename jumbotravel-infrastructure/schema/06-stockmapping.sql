CREATE DATABASE IF NOT EXISTS jumbotravel;
USE jumbotravel;

CREATE TABLE IF NOT EXISTS airplane_stockmapping (

    -- PRIMARY KEY
    stock_mappingid     INT NOT NULL AUTO_INCREMENT,
    stock_id            INT NOT NULL,

    -- FLAGS (for future use)
    created_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    -- CONSTRAINT
    PRIMARY KEY (stock_mappingid),
    FOREIGN KEY (stock_id) REFERENCES airplane_stock(stock_id)
);

