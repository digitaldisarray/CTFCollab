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
