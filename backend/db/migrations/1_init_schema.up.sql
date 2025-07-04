-- Users Table
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,  -- Adjust size based on hash algorithm
    is_admin BOOLEAN NOT NULL DEFAULT 0
);

-- CTFs Table
CREATE TABLE IF NOT EXISTS ctfs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    phrase VARCHAR(255) NOT NULL, -- Phrase used for joining CTF
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    start_date DATETIME NOT NULL,
    end_date DATETIME NOT NULL,
    author_id INT NOT NULL,  -- User who created the CTF
    FOREIGN KEY (author_id) REFERENCES users(id)
);

-- Guests Table (Anonymous users w/ nicknames)
CREATE TABLE IF NOT EXISTS guests (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nickname VARCHAR(255) NOT NULL,
    last_seen DATETIME DEFAULT CURRENT_TIMESTAMP,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    ctf_id INT NOT NULL,
    FOREIGN KEY (ctf_id) REFERENCES ctfs(id) ON DELETE CASCADE
);

-- User CTFs Table many-to-many
CREATE TABLE IF NOT EXISTS user_ctfs (
    user_id INT NOT NULL,
    ctf_id INT NOT NULL,
    joined_at DATETIME DEFAULT CURRENT_TIMESTAMP,  -- Timestamp of when the user joined the CTF
    PRIMARY KEY (user_id, ctf_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (ctf_id) REFERENCES ctfs(id) ON DELETE CASCADE
);

-- Challenges Table
CREATE TABLE IF NOT EXISTS challenges (
    id INT AUTO_INCREMENT PRIMARY KEY,
    ctf_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    flag TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    hedgedoc_url VARCHAR(512) UNIQUE NOT NULL,
    FOREIGN KEY (ctf_id) REFERENCES ctfs(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS ctf_participants (
    id INT AUTO_INCREMENT PRIMARY KEY,
    ctf_id INT NOT NULL,
    user_id INT DEFAULT NULL,
    guest_id INT DEFAULT NULL,
    joined_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY unique_user_per_ctf (ctf_id, user_id),
    UNIQUE KEY unique_guest_per_ctf (ctf_id, guest_id),
    CONSTRAINT chk_only_one_id CHECK (
        (user_id IS NULL AND guest_id IS NOT NULL) OR
        (user_id IS NOT NULL AND guest_id IS NULL)
    ),
    FOREIGN KEY (ctf_id) REFERENCES ctfs(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (guest_id) REFERENCES guests(id) ON DELETE CASCADE
);
-- -- Tags Table: Stores the different tags/categories users can assign to challenges
-- CREATE TABLE IF NOT EXISTS tags (
--     id INT AUTO_INCREMENT PRIMARY KEY,
--     name VARCHAR(255) UNIQUE NOT NULL  -- Name of the tag (e.g., "cryptography", "web", "forensics")
-- );

-- -- Challenge Tags Table: Many-to-many relationship between challenges and tags
-- CREATE TABLE IF NOT EXISTS challenge_tags (
--     challenge_id INT NOT NULL,
--     tag_id INT NOT NULL,
--     FOREIGN KEY (challenge_id) REFERENCES challenges(id) ON DELETE CASCADE,
--     FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE,
--     PRIMARY KEY (challenge_id, tag_id)  -- Ensures uniqueness of challenge-tag pairs
-- );

-- Notes Table
-- Might not need this if we use a pre existing tool for notes
-- CREATE TABLE IF NOT EXISTS notes (
--     id INT AUTO_INCREMENT PRIMARY KEY,
--     challenge_id INT NOT NULL,  -- Foreign key to challenge
--     created_by INT,             -- Foreign key to users (created by)
--     FOREIGN KEY (challenge_id) REFERENCES challenges(id),
--     FOREIGN KEY (created_by) REFERENCES users(id)
-- );
