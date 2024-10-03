CREATE TABLE IF NOT EXISTS recipes (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name TEXT NOT NULL UNIQUE,
    instruction LONGTEXT NOT NULL,
    added_by INT,
    FOREIGN KEY (added_by) references users(id)
)