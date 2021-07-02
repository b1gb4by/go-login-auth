DROP DATABASE IF EXISTS go_auth;
CREATE DATABASE IF NOT EXISTS go_auth;
USE go_auth;

DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users (
    id          INT(20)         NOT NULL AUTO_INCREMENT,
    first_name  VARCHAR(255)    NOT NULL,
    last_name   VARCHAR(255)    NOT NULL,
    email       VARCHAR(255)    NOT NULL,
    password    LONGTEXT        NOT NULL,
    created_at  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `email` (`email`)
);
