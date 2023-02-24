DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id SERIAL,
    username VARCHAR(128) UNIQUE,
    email VARCHAR(128) UNIQUE,
    passwordHash VARCHAR(32),
    PRIMARY KEY (id)
);