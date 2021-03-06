CREATE DATABASE IF NOT EXISTS jumbotravel;
USE jumbotravel;

CREATE TABLE IF NOT EXISTS notifications (

    notification_id             INT NOT NULL AUTO_INCREMENT,

    -- Scope (agent/group/flight/global)
    scope                       VARCHAR(255) NOT NULL,
    -- Can be used to link to resource (if provided must match with scope)
    resource_id                 INT NULL,
    -- Can be used to link to resource (if provided must match with scope)
    resource_uuid               VARCHAR(255) NULL,

    -- Notification
    title                       VARCHAR(255) NOT NULL,
    message                     VARCHAR(255) NULL,

    -- Link
    link                        TEXT NULL,
    extra                       TEXT NULL,

    -- Severity
    type                        VARCHAR(255) NOT NULL,

    -- Shown
    popup                       BOOLEAN NOT NULL DEFAULT FALSE,

    -- Dates
    expires_at                  TIMESTAMP NOT NULL,
    created_at                  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Flags
    seen                        BOOLEAN NOT NULL DEFAULT 0,
    active                      BOOLEAN NOT NULL DEFAULT 1,

    PRIMARY KEY (notification_id),
    INDEX resource_id_index (resource_id),
    INDEX resource_uuid_index (resource_uuid)
);