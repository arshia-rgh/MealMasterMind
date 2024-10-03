CREATE TABLE IF NOT EXISTS users
(
    id           INT PRIMARY KEY AUTO_INCREMENT,
    first_name   VARCHAR(255) NULL,
    last_name    VARCHAR(255) NULL,
    username     VARCHAR(255) UNIQUE,
    email        VARCHAR(255) UNIQUE,
    password     VARCHAR(255),
    phone_number VARCHAR(255) UNIQUE
);

CREATE TABLE IF NOT EXISTS recipes
(
    id          INT PRIMARY KEY AUTO_INCREMENT,
    name        VARCHAR(255) NOT NULL UNIQUE,
    instruction LONGTEXT     NOT NULL,
    added_by    INT,
    FOREIGN KEY (added_by) references users (id)
);

CREATE TABLE IF NOT EXISTS tags
(
    id        INT PRIMARY KEY AUTO_INCREMENT,
    name      VARCHAR(255) NOT NULL UNIQUE,
    recipe_id INT,
    FOREIGN KEY (recipe_id) REFERENCES recipes (id)
);

CREATE TABLE IF NOT EXISTS ingredients
(
    id        INT PRIMARY KEY AUTO_INCREMENT,
    name      VARCHAR(255) NOT NULL UNIQUE,
    recipe_id INT,
    FOREIGN KEY (recipe_id) REFERENCES recipes (id)
);

CREATE TABLE IF NOT EXISTS dietary_restrictions
(
    id   INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS user_dietary_restrictions
(
    user_id        INT,
    restriction_id INT,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (restriction_id) REFERENCES dietary_restrictions (id),
    PRIMARY KEY (user_id, restriction_id)
);

CREATE TABLE IF NOT EXISTS meal_plans
(
    id      INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    name    VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users (id)

);

CREATE TABLE IF NOT EXISTS meals
(
    id           INT PRIMARY KEY AUTO_INCREMENT,
    day          VARCHAR(100) NOT NULL,
    recipe_id    INT,
    meal_plan_id INT,
    FOREIGN KEY (recipe_id) REFERENCES recipes (id),
    FOREIGN KEY (meal_plan_id) REFERENCES meal_plans (id)
);

