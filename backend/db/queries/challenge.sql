-- name: GetChallenge :one
SELECT * FROM challenges
WHERE id = ? LIMIT 1;

-- name: SetFlag :execresult
UPDATE challenges
SET flag = ?
WHERE id = ?;

-- name: GetFlag :one
SELECT flag FROM challenges
WHERE id = ? LIMIT 1;

-- name: CreateChallenge :execresult
INSERT INTO challenges (ctf_id, name, description, flag, created_at)
SELECT 
    ctfs.id,
    ? AS challenge_name,        -- Replace ? with the challenge name
    ? AS challenge_description, -- Replace ? with the challenge description
    ? AS challenge_flag,        -- Replace ? with the challenge flag
    CURRENT_TIMESTAMP
FROM 
    ctfs
WHERE 
    ctfs.phrase = ?;

-- name: DeleteChallenge :exec
DELETE FROM challenges
WHERE id = ?;

-- name: UpdateChallenge :execresult
UPDATE challenges
SET
    name = COALESCE(sqlc.narg('name'), name),
    description = COALESCE(sqlc.narg('description'), description),
    flag = COALESCE(sqlc.narg('flag'), flag)
WHERE id = ?;
