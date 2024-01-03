CREATE TABLE
    todos (
        id INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
        user_id INT(11) UNSIGNED NOT NULL,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        is_done TINYINT NOT NULL DEFAULT 0,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY(id),
        FOREIGN KEY (user_id) REFERENCES users(id)
    ) Engine = InnoDb;