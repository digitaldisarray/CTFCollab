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
