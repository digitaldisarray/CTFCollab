-- name: CreateSession :execresult
INSERT INTO sessions (
    cookie, nickname
) VALUES (
    ?, ?
);

-- name: SeenSession :execresult
UPDATE sessions
SET last_seen = CURRENT_TIMESTAMP
WHERE cookie = ?;

-- name: GetSession :one
SELECT * FROM sessions
WHERE cookie = ? LIMIT 1;

-- name: PruneOldSessions :execresult
DELETE FROM sessions
WHERE last_seen < DATE_SUB(NOW(), INTERVAL 3 MONTH);
