CREATE TABLE users (
    id SERIAL,
    username VARCHAR(128),
    email VARCHAR(128) UNIQUE,
    passwordHash VARCHAR(32),
    PRIMARY KEY (id),
);