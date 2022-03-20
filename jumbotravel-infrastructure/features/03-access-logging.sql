CREATE DATABASE IF NOT EXISTS jumbotravel;
USE jumbotravel;

CREATE TABLE IF NOT EXISTS api_access (

    -- Request ID
    requestid                   VARCHAR(255) NOT NULL,
    
    -- Token Info
    token_id                    VARCHAR(255) NULL,
    token_name                  VARCHAR(255) NULL,

    -- Request Info
    ip                          VARCHAR(255) NULL,
    method                      VARCHAR(255) NOT NULL,
    path                        VARCHAR(255) NOT NULL,
    query                       VARCHAR(255) NULL,

    -- Response Info
    status                      INT NOT NULL,
    error_message               VARCHAR(255) NULL,

    -- Row Meta
    created_at                  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (requestid)
);