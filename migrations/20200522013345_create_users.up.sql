CREATE TABLE users (
    id bigserial NOT NULL PRIMARY KEY,
    email varchar NOT NULL unique,
    encrypted_password varchar NOT NULL
);