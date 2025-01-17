-- name: CreateUser :execresult
INSERT INTO users (
    username, password_hash
) VALUES (
    ?, ?
);

-- name: GetUser :one
SELECT * FROM users
WHERE username = ? LIMIT 1;

-- name: DeleteUser :execresult
DELETE from users
WHERE id = ?;

-- name: ChangePassword :execresult
UPDATE users
SET password_hash = ?
WHERE id = ?;

-- name: ChangeUsername :execresult
UPDATE users
SET username = ?
WHERE id = ?;
