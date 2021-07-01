DROP DATABASE IF EXISTS go_auth;
CREATE DATABASE IF NOT EXISTS go_auth;
USE go_auth;

DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users (
    id          INT(20)         NOT NULL AUTO_INCREMENT,
    first_name  VARCHAR(255)    NOT NULL,
    last_name   VARCHAR(255)    NOT NULL,
    email       VARCHAR(255)    NOT NULL,
    password    VARCHAR(255)    NOT NULL,
    created_at  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
);

INSERT INTO
    users (first_name, last_name, email, password)
VALUES
    (
        "Alan",
        "Bob",
        "alan@example.com",
        "test_password"
    );

INSERT INTO
    users (first_name, last_name, email, password)
VALUES
    (
        "Callen",
        "Den",
        "callen@example.com",
        "test_password1"
    );
