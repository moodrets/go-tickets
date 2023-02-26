-- Active: 1677391540282@@127.0.0.1@3310@tickets
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS tickets;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    login VARCHAR(255) UNIQUE,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE tickets (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    creator_id BIGINT,
    in_work_user_id BIGINT UNSIGNED,
    title VARCHAR(255),
    status ENUM('new', 'in_work', 'wait_answer', 'closed') DEFAULT 'new',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT tickets_users_fk FOREIGN KEY (in_work_user_id) REFERENCES users (id) ON DELETE CASCADE
);
CREATE TABLE comments (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    ticket_id BIGINT UNSIGNED,
    author_id BIGINT UNSIGNED,
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT comments_ticket_fk FOREIGN KEY (ticket_id) REFERENCES tickets (id) ON DELETE CASCADE,
    CONSTRAINT comments_users_fk FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE CASCADE
);

-- add users
INSERT INTO users (login, name) VALUES
    ('Dimon', 'Dima'),
    ('Geniy', 'Evgeniy'),
    ('Teamlead', 'Maksim'),
    ('Paul', 'Pavel');
-- add tickets
INSERT INTO tickets (creator_id, in_work_user_id, title, status) VALUES 
    (1, 1, "Тикет 1", "new"),
    (1, 2, "Тикет 2", "in_work"),
    (1, 3, "Тикет 3", "wait_answer"),
    (2, 4, "Тикет 4", "new");
-- add comments
INSERT INTO comments (ticket_id, author_id, comment) VALUES 
    (1, 1, "Комментарий 1"),
    (1, 2, "Комментарий 2"),
    (2, 2, "Комментарий 3"),
    (2, 4, "Комментарий 4");