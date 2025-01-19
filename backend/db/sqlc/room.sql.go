// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: room.sql

package db

import (
	"context"
	"database/sql"
)

const createRoom = `-- name: CreateRoom :execresult
INSERT INTO rooms (
    room_id, name, description
) VALUES (
    ?, ?, ?
)
`

type CreateRoomParams struct {
	RoomID      int32          `json:"room_id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) CreateRoom(ctx context.Context, arg CreateRoomParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createRoom, arg.RoomID, arg.Name, arg.Description)
}

const deleteRoom = `-- name: DeleteRoom :execresult
DELETE FROM rooms
WHERE room_id = ?
`

func (q *Queries) DeleteRoom(ctx context.Context, roomID int32) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteRoom, roomID)
}

const getAllRooms = `-- name: GetAllRooms :many
SELECT id, room_id, name, description, created_at, updated_at FROM rooms
`

func (q *Queries) GetAllRooms(ctx context.Context) ([]Room, error) {
	rows, err := q.db.QueryContext(ctx, getAllRooms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Room
	for rows.Next() {
		var i Room
		if err := rows.Scan(
			&i.ID,
			&i.RoomID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRoom = `-- name: GetRoom :one
SELECT id, room_id, name, description, created_at, updated_at FROM rooms
WHERE room_id = ? LIMIT 1
`

func (q *Queries) GetRoom(ctx context.Context, roomID int32) (Room, error) {
	row := q.db.QueryRowContext(ctx, getRoom, roomID)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.RoomID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}