-- name: CreateRoom :execresult
INSERT INTO rooms (
    room_id, name, description
) VALUES (
    ?, ?, ?
);

-- name: GetRoom :one
SELECT * FROM rooms
WHERE room_id = ? LIMIT 1;

-- name: GetAllRooms :many
SELECT * FROM rooms;

-- name: DeleteRoom :execresult
DELETE FROM rooms
WHERE room_id = ?;
