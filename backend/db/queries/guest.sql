-- name: CreateGuest :execresult
INSERT INTO guests (
    nickname
) VALUES (
    ?
);

-- name: MarkGuestSeen :execresult
UPDATE guests
SET last_seen = CURRENT_TIMESTAMP
WHERE nickname = ?;

-- name: GetGuestByID :one
SELECT * FROM guests
WHERE id = ? LIMIT 1;

-- name: GetGuestByName :one
SELECT * FROM guests
WHERE nickname = ? LIMIT 1;

-- name: PruneOldSessions :execresult
DELETE FROM guests
WHERE last_seen < DATE_SUB(NOW(), INTERVAL 3 MONTH);
