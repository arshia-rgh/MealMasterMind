CREATE TABLE  IF NOT EXISTS users(
    id INT PRIMARY KEY AUTO_INCREMENT,
    first_name VARCHAR(255) NULL,
    last_name VARCHAR(255) NULL,
    username VARCHAR(255) UNIQUE,
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    phone_number VARCHAR(255) UNIQUE
);