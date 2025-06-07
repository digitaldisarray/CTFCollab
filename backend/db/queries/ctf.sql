-- name: ListAllCTFs :many
SELECT * FROM ctfs
ORDER BY start_date;

-- name: SearchCTFs :many
SELECT
    *
FROM
    ctfs
WHERE
    (sqlc.arg(name) IS NULL OR ctfs.name LIKE CONCAT('%', sqlc.arg(name), '%')) 
    and (sqlc.arg(description) IS NULL OR ctfs.description LIKE CONCAT('%', sqlc.arg(description), '%'))
ORDER BY start_date;

-- name: GetCTFByPhrase :one
SELECT * FROM ctfs
WHERE phrase = ? LIMIT 1;

-- name: CreateCTF :execresult
INSERT INTO ctfs (
    name, phrase, description, start_date, end_date, author_id
) VALUES (
    ?, ?, ?, ?, ?, ?
);

-- name: UpdateCTF :execresult
UPDATE ctfs
SET
    name = ?,
    description = ?,
    start_date = ?,
    end_date = ?
WHERE phrase = ?;

-- name: DeleteCTFByPhrase :exec
DELETE FROM ctfs
WHERE phrase = ?;

-- name: JoinCTFUser :execresult
INSERT INTO user_ctfs (
    user_id, ctf_id
) VALUES (
    ?, ?
);

-- name: ListUsersJoinedCTFs :many
SELECT 
    ctfs.name AS ctf_name,
    ctfs.description AS ctf_description,
    ctfs.start_date,
    ctfs.end_date,
    ctfs.author_id AS ctf_author_id,
    ctfs.phrase
FROM 
    ctfs
JOIN 
    user_ctfs ON ctfs.id = user_ctfs.ctf_id
WHERE 
    user_ctfs.user_id = ?;

-- name: ListGuestsJoinedCTF :many
SELECT ctf_id FROM guests
WHERE id = ?;

-- name: GetCTFChallenges :many
SELECT 
    challenges.id AS challenge_id,
    challenges.name AS challenge_name,
    challenges.description AS challenge_description,
    challenges.flag,
    challenges.created_at AS challenge_created_at,
    challenges.hedgedoc_url as hedgedoc_url
FROM 
    challenges
JOIN 
    ctfs ON challenges.ctf_id = ctfs.id
WHERE 
    ctfs.phrase = ?;

-- name: IsUserMemberOfCTF :one
SELECT EXISTS (
    SELECT 1
    FROM user_ctfs
    JOIN ctfs ON user_ctfs.ctf_id = ctfs.id
    WHERE user_ctfs.user_id = ? AND ctfs.phrase = ?
) AS is_member;

-- name: IsGuestMemberOfCTF :one
SELECT EXISTS (
    SELECT 1
    FROM guests
    JOIN ctfs ON guests.ctf_id = ctfs.id
    WHERE guests.id = ? AND ctfs.phrase = ?
) AS is_member;

-- name: InsertParticipant :exec
INSERT INTO ctf_participants (ctf_id, user_id, guest_id)
VALUES (?, ?, ?)
ON DUPLICATE KEY UPDATE joined_at = CURRENT_TIMESTAMP;

-- name: RemoveParticipantByUser :exec
DELETE FROM ctf_participants
WHERE ctf_id = ? AND user_id = ?;

-- name: RemoveParticipantByGuest :exec
DELETE FROM ctf_participants
WHERE ctf_id = ? AND guest_id = ?;

-- name: GetParticipantsByCTF :many
SELECT
    cp.id,
    cp.ctf_id,
    cp.user_id,
    cp.guest_id,
    u.username,
    g.nickname
FROM ctf_participants cp
LEFT JOIN users u ON cp.user_id = u.id
LEFT JOIN guests g ON cp.guest_id = g.id
WHERE cp.ctf_id = ?;