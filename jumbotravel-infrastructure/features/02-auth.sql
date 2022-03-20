CREATE DATABASE IF NOT EXISTS jumbotravel;
USE jumbotravel;

CREATE TABLE IF NOT EXISTS agent_auth (

    -- Token Data
    id                      VARCHAR(255) NOT NULL,
    subjet                  VARCHAR(255) NOT NULL,
    token                   TEXT NOT NULL,

    -- User Data (agentmapping id)
    agent_id                INT NOT NULL,

    -- Flags
    issued_at               TIMESTAMP NOT NULL,
    expires_at              TIMESTAMP NOT NULL,

    active                  BOOLEAN NOT NULL DEFAULT 1,

    -- Row Meta
    created_at              TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (agent_id) REFERENCES master_agentmapping(agentmapping_id)
);