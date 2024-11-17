-- name: CreateUser :one
INSERT INTO users (
    username, password_hash
) VALUES (
    ?, ?
)
RETURNING *;

-- name: DeleteUser :exec
DELETE from users
WHERE id = ?;

-- name: ChangePassword :exec
UPDATE users
SET password_hash = ?
WHERE id = ?;

-- name: ChangeUsername :exec
UPDATE users
SET username = ?
WHERE id = ?;



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



-- name: ListCTFs :many
SELECT * FROM ctfs
ORDER BY start_date;

-- name: CreateCTF :one
INSERT INTO ctfs (
    name, description, start_date, end_date, author_id
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateCTF :exec
UPDATE ctfs
SET
    name = COALESCE(sqlc.narg('name'), name),
    description = COALESCE(sqlc.narg('description'), description),
    start_date = COALESCE(sqlc.narg('start_date'), start_date),
    end_date = COALESCE(sqlc.narg('end_date'), end_date)
WHERE id = ?;

-- name: DeleteCTF :exec
DELETE FROM ctfs
WHERE id = ?;



-- name: ListChallenges :many
SELECT * FROM challenges
ORDER BY created_at;

-- name: GetChallenge :one
SELECT * FROM challenges
WHERE id = ? LIMIT 1;

-- name: SetFlag :exec
UPDATE challenges
SET flag = ?
WHERE id = ?;

-- name: GetFlag :one
SELECT flag FROM challenges
WHERE id = ? LIMIT 1;

-- name: CreateChallenge :one
INSERT INTO challenges (
    ctf_id, name, description
) VALUES (
    ?, ?, ?
)
RETURNING *;

-- name: DeleteChallenge :exec
BEGIN TRANSACTION;
-- Delete entries from challenge_tags to remove associations
DELETE FROM challenge_tags WHERE challenge_id = ?;
-- Delete the challenge
DELETE FROM challenges WHERE id = ?;
-- Delete tags that are no longer associated with any challenges
DELETE FROM tags WHERE id NOT IN (SELECT DISTINCT tag_id FROM challenge_tags);
COMMIT;

-- name: UpdateChallenge :exec
UPDATE challenges
SET
    name = COALESCE(sqlc.narg('name'), name),
    description = COALESCE(sqlc.narg('description'), description),
    flag = COALESCE(sqlc.narg('flag'), flag)
WHERE id = ?;

-- name: GetChallengeCategories :many
SELECT t.id, t.name
FROM tags t
JOIN challenge_tags ct ON t.id = ct.tag_id
WHERE ct.challenge_id = ?;

-- name: AddTag :exec
BEGIN TRANSACTION;
-- Step 1: Insert the tag if it doesn't already exist
INSERT INTO tags (name)
VALUES (?)
ON CONFLICT(name) DO NOTHING;
-- Step 2: Get the tag ID for the inserted or existing tag
SELECT id FROM tags WHERE name = ?;
-- Step 3: Insert the challenge-tag relationship if it doesn't exist
INSERT INTO challenge_tags (challenge_id, tag_id)
SELECT ?, id FROM tags WHERE name = ? 
ON CONFLICT(challenge_id, tag_id) DO NOTHING;
COMMIT;

-- name: DeleteTag :exec
BEGIN TRANSACTION;
-- Step 1: Delete challenge-tag relationships for the given tag
DELETE FROM challenge_tags WHERE tag_id = ?;
-- Step 2: Delete the tag itself if it's no longer in use (no other references in challenge_tags)
DELETE FROM tags WHERE id = ? AND NOT EXISTS (
    SELECT 1 FROM challenge_tags WHERE tag_id = tags.id
);
COMMIT;
