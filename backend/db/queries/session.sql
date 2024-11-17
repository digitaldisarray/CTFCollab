-- name: CreateSession :one
INSERT INTO sessions (
    cookie, nickname
) VALUES (
    ?, ?
)
RETURNING *;

-- name: SeenSession :exec
UPDATE sessions
SET last_seen = CURRENT_TIMESTAMP
WHERE cookie = ?;

-- name: GetSession :one
SELECT * FROM sessions
WHERE cookie = ? LIMIT 1;

-- name: DeleteOldSessions :one
DELETE FROM sessions
WHERE last_seen < DATETIME('now', '-3 months')
RETURNING *;
