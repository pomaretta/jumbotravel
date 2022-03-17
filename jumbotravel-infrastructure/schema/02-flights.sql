CREATE DATABASE IF NOT EXISTS jumbotravel;
USE jumbotravel;

-- FLIGHTS TABLES
CREATE TABLE IF NOT EXISTS flight_routes (

    route_id                INT NOT NULL AUTO_INCREMENT,
    airplane_id             INT NOT NULL,

    departure_country       VARCHAR(255) NOT NULL,    
    arrival_country         VARCHAR(255) NOT NULL,
    departure_city          VARCHAR(255) NOT NULL,
    arrival_city            VARCHAR(255) NOT NULL,
    departure_airport       VARCHAR(255) NOT NULL,
    arrival_airport         VARCHAR(255) NOT NULL,

    created_at              TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (route_id),
    FOREIGN KEY (airplane_id) REFERENCES master_airplanes(airplane_id)
);
CREATE TABLE IF NOT EXISTS flights (

    flight_id               INT NOT NULL AUTO_INCREMENT,
    route_id                INT NOT NULL,
    status                  VARCHAR(255) NOT NULL,

    departure_country       VARCHAR(255) NOT NULL,    
    arrival_country         VARCHAR(255) NOT NULL,
    departure_city          VARCHAR(255) NOT NULL,
    arrival_city            VARCHAR(255) NOT NULL,
    departure_airport       VARCHAR(255) NOT NULL,
    arrival_airport         VARCHAR(255) NOT NULL,

    departure_date          TIMESTAMP NOT NULL,
    arrival_date            TIMESTAMP NOT NULL,

    created_at              TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (flight_id),
    FOREIGN KEY (route_id) REFERENCES flight_routes(route_id)
);
CREATE TABLE IF NOT EXISTS flight_agents (

    flight_id               INT NOT NULL,
    agent_id                INT NOT NULL,

    PRIMARY KEY (flight_id, agent_id),
    FOREIGN KEY (flight_id) REFERENCES flights(flight_id),
    FOREIGN KEY (agent_id) REFERENCES master_agents(agent_id)
);
