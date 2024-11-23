-- name: ListCTFs :many
SELECT * FROM ctfs
ORDER BY start_date;

-- name: CreateCTF :execresult
INSERT INTO ctfs (
    name, description, start_date, end_date, author_id
) VALUES (
    ?, ?, ?, ?, ?
);

-- name: UpdateCTF :execresult
UPDATE ctfs
SET
    name = ?,
    description = ?,
    start_date = ?,
    end_date = ?
WHERE id = ?;

-- name: DeleteCTF :exec
DELETE FROM ctfs
WHERE id = ?;
