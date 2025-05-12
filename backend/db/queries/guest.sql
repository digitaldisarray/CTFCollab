-- name: CreateGuest :execresult
INSERT INTO guests (
    nickname,
    ctf_id
) VALUES (
    ?, ?
);

-- name: MarkGuestSeen :execresult
UPDATE guests
SET last_seen = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: GetGuestByID :one
SELECT * FROM guests
WHERE id = ? LIMIT 1;

-- name: GetGuestByNameAndCTF :one
SELECT * FROM guests
WHERE nickname = ? AND ctf_id = ? LIMIT 1;

-- name: PruneOldSessions :execresult
DELETE FROM guests
WHERE last_seen < DATE_SUB(NOW(), INTERVAL 3 MONTH);
