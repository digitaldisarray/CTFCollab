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

-- name: JoinCTFGuest :execresult
INSERT INTO guest_ctfs (
    guest_id, ctf_id
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

-- name: ListGuestsJoinedCTFs :many
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
    guest_ctfs ON ctfs.id = guest_ctfs.ctf_id
WHERE 
    guest_ctfs.guest_id = ?;


-- name: GetCTFChallenges :many
SELECT 
    challenges.id AS challenge_id,
    challenges.name AS challenge_name,
    challenges.description AS challenge_description,
    challenges.flag,
    challenges.created_at AS challenge_created_at
FROM 
    challenges
JOIN 
    ctfs ON challenges.ctf_id = ctfs.id
WHERE 
    ctfs.phrase = ?;

-- name: IsUserMemberOfCTF :one
SELECT COUNT(*) > 0 AS is_member
FROM user_ctfs
JOIN ctfs ON user_ctfs.ctf_id = ctfs.id
WHERE user_ctfs.user_id = ? AND ctfs.phrase = ?;

-- name: IsGuestMemberOfCTF :one
SELECT COUNT(*) > 0 AS is_member
FROM guest_ctfs
JOIN ctfs ON guest_ctfs.ctf_id = ctfs.id
WHERE guest_ctfs.guest_id = ? AND ctfs.phrase = ?;
