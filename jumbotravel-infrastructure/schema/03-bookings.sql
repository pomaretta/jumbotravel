CREATE DATABASE IF NOT EXISTS jumbotravel;
USE jumbotravel;

-- BOOKING TABLES
CREATE TABLE IF NOT EXISTS bookings (

    bookingreferenceid      VARCHAR(255) NOT NULL,
    bookingitemid           VARCHAR(255) NOT NULL,
    status                  VARCHAR(255) NOT NULL,

    agent_id                INT NOT NULL,
    product_id              INT NOT NULL,
    flight_id               INT NOT NULL,

    items                   INT NOT NULL,
    price                   FLOAT NOT NULL,

    -- PROVIDER
    provider_id             INT NULL,

    -- FLAGS (for future use)
    created_at              TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    -- INTEGRITY
    hash64                  INT NOT NULL,

    -- CONSTRAINT
    PRIMARY KEY (bookingreferenceid, bookingitemid),
    FOREIGN KEY (agent_id) REFERENCES master_agents(agent_id),
    FOREIGN KEY (product_id) REFERENCES master_products(product_id),
    FOREIGN KEY (flight_id) REFERENCES flights(flight_id)
);