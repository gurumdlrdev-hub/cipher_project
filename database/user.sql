CREATE TABLE app_user (
User_id SERIAL,
username VARCHAR(50) UNIQUE NOT NULL,
password_hash VARCHAR(250) NOT NULL
);

SELECT * FROM app_user;

DROP TABLE IF EXISTS app_user;

CREATE TABLE app_user (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);