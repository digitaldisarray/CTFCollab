-- Users Table (Admins only for now)
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT
);

-- Sessions Table (Anonymous users w/ nicknames)
CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    cookie TEXT UNIQUE NOT NULL,
    nickname TEXT NOT NULL,
    last_seen DATETIME DEFAULT CURRENT_TIMESTAMP,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- CTFs Table
CREATE TABLE IF NOT EXISTS ctfs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT,
    start_date DATETIME,
    end_date DATETIME,
    author_id INTEGER,  -- Admin user who created the CTF
    FOREIGN KEY (author_id) REFERENCES users(id)
);

-- Challenges Table
CREATE TABLE IF NOT EXISTS challenges (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    ctf_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    flag TEXT,
    note_id INTEGER,  -- One-to-one relationship with notes
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (ctf_id) REFERENCES ctfs(id),
    FOREIGN KEY (note_id) REFERENCES notes(id) ON DELETE CASCADE
);

-- Tags Table: Stores the different tags (categories) users can assign to challenges
CREATE TABLE IF NOT EXISTS tags (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE NOT NULL  -- Name of the tag (e.g., "cryptography", "web", "forensics")
);

-- Challenge Tags Table: Many-to-many relationship between challenges and tags
CREATE TABLE IF NOT EXISTS challenge_tags (
    challenge_id INTEGER NOT NULL,
    tag_id INTEGER NOT NULL,
    FOREIGN KEY (challenge_id) REFERENCES challenges(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (challenge_id, tag_id)  -- Ensures uniqueness of challenge-tag pairs
);

-- Notes Table
CREATE TABLE IF NOT EXISTS notes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    challenge_id INTEGER NOT NULL,  -- Foreign key to challenge
    -- More data needed, maybe peers or a link to embed an editor
    FOREIGN KEY (challenge_id) REFERENCES challenges(id),
    FOREIGN KEY (created_by) REFERENCES users(id)
);
