// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const changePassword = `-- name: ChangePassword :execresult
UPDATE users
SET password_hash = ?
WHERE id = ?
`

type ChangePasswordParams struct {
	PasswordHash string `json:"password_hash"`
	ID           int32  `json:"id"`
}

func (q *Queries) ChangePassword(ctx context.Context, arg ChangePasswordParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, changePassword, arg.PasswordHash, arg.ID)
}

const changeUsername = `-- name: ChangeUsername :execresult
UPDATE users
SET username = ?
WHERE id = ?
`

type ChangeUsernameParams struct {
	Username string `json:"username"`
	ID       int32  `json:"id"`
}

func (q *Queries) ChangeUsername(ctx context.Context, arg ChangeUsernameParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, changeUsername, arg.Username, arg.ID)
}

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (
    username, password_hash
) VALUES (
    ?, ?
)
`

type CreateUserParams struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser, arg.Username, arg.PasswordHash)
}

const deleteUser = `-- name: DeleteUser :execresult
DELETE from users
WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteUser, id)
}

const getUser = `-- name: GetUser :one
SELECT id, username, password_hash FROM users
WHERE username = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.PasswordHash)
	return i, err
}
