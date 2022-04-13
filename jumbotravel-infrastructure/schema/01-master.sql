CREATE DATABASE IF NOT EXISTS jumbotravel;
USE jumbotravel;

-- MASTER TABLES
CREATE TABLE IF NOT EXISTS master_airplanes (

    -- PRIMARY KEY
    airplane_id         INT NOT NULL AUTO_INCREMENT,

    -- METADATA
    carrier             VARCHAR(255) NOT NULL,
    flight_number       VARCHAR(255) NOT NULL,
    seats               INT NOT NULL,

    -- FLAGS (for future use)
    created_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    -- CONSTRAINT
    PRIMARY KEY (airplane_id)
);
CREATE TABLE IF NOT EXISTS master_agents (

    -- PRIMARY KEY
    agent_id            INT NOT NULL AUTO_INCREMENT,
    -- UNIQUE KEY (DNI)
    dni                 VARCHAR(255) NULL,

    -- IDENTIFIERS (NAME, SURNAME, EMAIL)
    name                VARCHAR(255) NOT NULL,
    surname             VARCHAR(255) NOT NULL,
    email               VARCHAR(255) NOT NULL,

    -- PASSWORD
    password            VARCHAR(255) NOT NULL,

    -- TYPE OF USER
    type                VARCHAR(255) NOT NULL,
    -- IF TYPE IS PROVIDER THEN WE NEED TO STORE THE AIRPORT ID
    airport_id          INT NULL,

    -- CREATION DATE
    created_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- ACTIVE
    active              BOOLEAN DEFAULT TRUE,

    -- CONSTRAINT
    PRIMARY KEY (agent_id)
);
CREATE TABLE IF NOT EXISTS master_products (

    -- PRIMARY KEY
    product_id          INT NOT NULL AUTO_INCREMENT,
    product_code        INT NOT NULL,
    
    -- METADATA
    name                VARCHAR(255) NOT NULL,
    description         VARCHAR(255) NOT NULL,
    brand               VARCHAR(255) NOT NULL,
    type                VARCHAR(255) NOT NULL,
    max                 INT NOT NULL,

    -- PRICE
    saleprice           FLOAT NOT NULL,

    -- ACTIVE
    active              BOOLEAN NOT NULL DEFAULT TRUE,

    -- CREATION DATE
    created_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- CONSTRAINT
    PRIMARY KEY (product_id),
    UNIQUE (product_code)
);