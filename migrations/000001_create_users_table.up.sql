CREATE TABLE users (
    id SERIAL PRIMARY KEY, 
    name VARCHAR(255) NOT NULL,                 
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL , 
    country VARCHAR(255) NOT NULL ,
    occupation VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL
);