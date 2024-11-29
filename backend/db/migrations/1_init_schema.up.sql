-- Users Table (Admins only for now)
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL  -- Adjust size based on your hash algorithm
);

-- Sessions Table (Anonymous users w/ nicknames)
CREATE TABLE IF NOT EXISTS sessions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cookie VARCHAR(255) UNIQUE NOT NULL,
    nickname VARCHAR(255) NOT NULL,
    last_seen DATETIME DEFAULT CURRENT_TIMESTAMP,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- CTFs Table
CREATE TABLE IF NOT EXISTS ctfs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    start_date DATETIME NOT NULL,
    end_date DATETIME NOT NULL,
    author_id INT NOT NULL,  -- Admin user who created the CTF
    FOREIGN KEY (author_id) REFERENCES users(id)
);

-- Challenges Table
CREATE TABLE IF NOT EXISTS challenges (
    id INT AUTO_INCREMENT PRIMARY KEY,
    ctf_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    flag TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (ctf_id) REFERENCES ctfs(id)
);

-- Tags Table: Stores the different tags/categories users can assign to challenges
CREATE TABLE IF NOT EXISTS tags (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL  -- Name of the tag (e.g., "cryptography", "web", "forensics")
);

-- Challenge Tags Table: Many-to-many relationship between challenges and tags
CREATE TABLE IF NOT EXISTS challenge_tags (
    challenge_id INT NOT NULL,
    tag_id INT NOT NULL,
    FOREIGN KEY (challenge_id) REFERENCES challenges(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (challenge_id, tag_id)  -- Ensures uniqueness of challenge-tag pairs
);

-- Notes Table
-- Might not need this if we use a pre existing tool for notes
-- CREATE TABLE IF NOT EXISTS notes (
--     id INT AUTO_INCREMENT PRIMARY KEY,
--     challenge_id INT NOT NULL,  -- Foreign key to challenge
--     created_by INT,             -- Foreign key to users (created by)
--     FOREIGN KEY (challenge_id) REFERENCES challenges(id),
--     FOREIGN KEY (created_by) REFERENCES users(id)
-- );
