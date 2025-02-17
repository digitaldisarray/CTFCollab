-- name: CreateUser :execresult
INSERT INTO users (
    username, password_hash, is_admin
) VALUES (
    ?, ?, 0
);

-- name: CreateAdmin :execresult
INSERT INTO users (
    username, password_hash, is_admin
) VALUES (
    ?, ?, 1
);

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = ? LIMIT 1;

-- name: DeleteUser :execresult
DELETE from users
WHERE username = ?;

-- name: ChangePassword :execresult
UPDATE users
SET password_hash = ?
WHERE username = ?;

-- name: ChangeUsername :execresult
UPDATE users
SET username = ?
WHERE id = ?;

-- name: SetAdminStatus :execresult
UPDATE users
SET is_admin = ?
WHERE username = ?;
