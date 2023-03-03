CREATE TYPE billing_period AS ENUM ('monthly', 'quarterly', 'yearly');

DROP TABLE IF EXISTS subscription;
CREATE TABLE subscription (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    start_date DATE NOT NULL,
    billing_period billing_period NOT NULL,
    price NUMERIC(15, 6),
    user_id INTEGER NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id)
);