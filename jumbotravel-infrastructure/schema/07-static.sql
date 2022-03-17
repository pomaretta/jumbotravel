CREATE DATABASE IF NOT EXISTS jumbotravel;
USE jumbotravel;

CREATE TABLE IF NOT EXISTS master_airports (

    airport_id          INT NOT NULL AUTO_INCREMENT,

    country             VARCHAR(255) NOT NULL,
    city                VARCHAR(255) NOT NULL,
    airport             VARCHAR(255) NOT NULL,
    common_name         VARCHAR(255) NOT NULL,   

    created_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (airport_id)
);