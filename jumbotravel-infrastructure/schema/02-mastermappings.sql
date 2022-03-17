CREATE DATABASE IF NOT EXISTS jumbotravel;
USE jumbotravel;

-- Airplanes Mapping
CREATE TABLE IF NOT EXISTS master_airplanesmapping (

    airplanemapping_id        INT NOT NULL AUTO_INCREMENT,
    airplane_id               INT NOT NULL,

    created_at                TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at                TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (airplanemapping_id),
    FOREIGN KEY (airplane_id) REFERENCES master_airplanes(airplane_id)
);

CREATE TABLE IF NOT EXISTS master_agentmapping (

    agentmapping_id           INT NOT NULL AUTO_INCREMENT,
    agent_id                  INT NOT NULL,

    created_at                TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at                TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (agentmapping_id),
    FOREIGN KEY (agent_id) REFERENCES master_agents(agent_id)
);

CREATE TABLE IF NOT EXISTS master_productmapping (

    productmapping_id         INT NOT NULL AUTO_INCREMENT,
    product_id                INT NOT NULL,

    created_at                TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at                TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (productmapping_id),
    FOREIGN KEY (product_id) REFERENCES master_products(product_id)
);