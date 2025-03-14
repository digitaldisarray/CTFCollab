// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: ctf.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createCTF = `-- name: CreateCTF :execresult
INSERT INTO ctfs (
    name, phrase, description, start_date, end_date, author_id
) VALUES (
    ?, ?, ?, ?, ?, ?
)
`

type CreateCTFParams struct {
	Name        string    `json:"name"`
	Phrase      string    `json:"phrase"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	AuthorID    int32     `json:"author_id"`
}

func (q *Queries) CreateCTF(ctx context.Context, arg CreateCTFParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createCTF,
		arg.Name,
		arg.Phrase,
		arg.Description,
		arg.StartDate,
		arg.EndDate,
		arg.AuthorID,
	)
}

const deleteCTFByPhrase = `-- name: DeleteCTFByPhrase :exec
DELETE FROM ctfs
WHERE phrase = ?
`

func (q *Queries) DeleteCTFByPhrase(ctx context.Context, phrase string) error {
	_, err := q.db.ExecContext(ctx, deleteCTFByPhrase, phrase)
	return err
}

const getCTFByPhrase = `-- name: GetCTFByPhrase :one
SELECT id, phrase, name, description, start_date, end_date, author_id FROM ctfs
WHERE phrase = ? LIMIT 1
`

func (q *Queries) GetCTFByPhrase(ctx context.Context, phrase string) (Ctf, error) {
	row := q.db.QueryRowContext(ctx, getCTFByPhrase, phrase)
	var i Ctf
	err := row.Scan(
		&i.ID,
		&i.Phrase,
		&i.Name,
		&i.Description,
		&i.StartDate,
		&i.EndDate,
		&i.AuthorID,
	)
	return i, err
}

const getCTFChallenges = `-- name: GetCTFChallenges :many
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
    ctfs.phrase = ?
`

type GetCTFChallengesRow struct {
	ChallengeID          int32        `json:"challenge_id"`
	ChallengeName        string       `json:"challenge_name"`
	ChallengeDescription string       `json:"challenge_description"`
	Flag                 string       `json:"flag"`
	ChallengeCreatedAt   sql.NullTime `json:"challenge_created_at"`
	HedgedocUrl          string       `json:"hedgedoc_url"`
}

func (q *Queries) GetCTFChallenges(ctx context.Context, phrase string) ([]GetCTFChallengesRow, error) {
	rows, err := q.db.QueryContext(ctx, getCTFChallenges, phrase)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCTFChallengesRow
	for rows.Next() {
		var i GetCTFChallengesRow
		if err := rows.Scan(
			&i.ChallengeID,
			&i.ChallengeName,
			&i.ChallengeDescription,
			&i.Flag,
			&i.ChallengeCreatedAt,
			&i.HedgedocUrl,
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

const isGuestMemberOfCTF = `-- name: IsGuestMemberOfCTF :one
SELECT COUNT(*) > 0 AS is_member
FROM guest_ctfs
JOIN ctfs ON guest_ctfs.ctf_id = ctfs.id
WHERE guest_ctfs.guest_id = ? AND ctfs.phrase = ?
`

type IsGuestMemberOfCTFParams struct {
	GuestID int32  `json:"guest_id"`
	Phrase  string `json:"phrase"`
}

func (q *Queries) IsGuestMemberOfCTF(ctx context.Context, arg IsGuestMemberOfCTFParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, isGuestMemberOfCTF, arg.GuestID, arg.Phrase)
	var is_member bool
	err := row.Scan(&is_member)
	return is_member, err
}

const isUserMemberOfCTF = `-- name: IsUserMemberOfCTF :one
SELECT COUNT(*) > 0 AS is_member
FROM user_ctfs
JOIN ctfs ON user_ctfs.ctf_id = ctfs.id
WHERE user_ctfs.user_id = ? AND ctfs.phrase = ?
`

type IsUserMemberOfCTFParams struct {
	UserID int32  `json:"user_id"`
	Phrase string `json:"phrase"`
}

func (q *Queries) IsUserMemberOfCTF(ctx context.Context, arg IsUserMemberOfCTFParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, isUserMemberOfCTF, arg.UserID, arg.Phrase)
	var is_member bool
	err := row.Scan(&is_member)
	return is_member, err
}

const joinCTFGuest = `-- name: JoinCTFGuest :execresult
INSERT INTO guest_ctfs (
    guest_id, ctf_id
) VALUES (
    ?, ?
)
`

type JoinCTFGuestParams struct {
	GuestID int32 `json:"guest_id"`
	CtfID   int32 `json:"ctf_id"`
}

func (q *Queries) JoinCTFGuest(ctx context.Context, arg JoinCTFGuestParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, joinCTFGuest, arg.GuestID, arg.CtfID)
}

const joinCTFUser = `-- name: JoinCTFUser :execresult
INSERT INTO user_ctfs (
    user_id, ctf_id
) VALUES (
    ?, ?
)
`

type JoinCTFUserParams struct {
	UserID int32 `json:"user_id"`
	CtfID  int32 `json:"ctf_id"`
}

func (q *Queries) JoinCTFUser(ctx context.Context, arg JoinCTFUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, joinCTFUser, arg.UserID, arg.CtfID)
}

const listAllCTFs = `-- name: ListAllCTFs :many
SELECT id, phrase, name, description, start_date, end_date, author_id FROM ctfs
ORDER BY start_date
`

func (q *Queries) ListAllCTFs(ctx context.Context) ([]Ctf, error) {
	rows, err := q.db.QueryContext(ctx, listAllCTFs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Ctf
	for rows.Next() {
		var i Ctf
		if err := rows.Scan(
			&i.ID,
			&i.Phrase,
			&i.Name,
			&i.Description,
			&i.StartDate,
			&i.EndDate,
			&i.AuthorID,
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

const listGuestsJoinedCTFs = `-- name: ListGuestsJoinedCTFs :many
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
    guest_ctfs.guest_id = ?
`

type ListGuestsJoinedCTFsRow struct {
	CtfName        string    `json:"ctf_name"`
	CtfDescription string    `json:"ctf_description"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	CtfAuthorID    int32     `json:"ctf_author_id"`
	Phrase         string    `json:"phrase"`
}

func (q *Queries) ListGuestsJoinedCTFs(ctx context.Context, guestID int32) ([]ListGuestsJoinedCTFsRow, error) {
	rows, err := q.db.QueryContext(ctx, listGuestsJoinedCTFs, guestID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListGuestsJoinedCTFsRow
	for rows.Next() {
		var i ListGuestsJoinedCTFsRow
		if err := rows.Scan(
			&i.CtfName,
			&i.CtfDescription,
			&i.StartDate,
			&i.EndDate,
			&i.CtfAuthorID,
			&i.Phrase,
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

const listUsersJoinedCTFs = `-- name: ListUsersJoinedCTFs :many
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
    user_ctfs.user_id = ?
`

type ListUsersJoinedCTFsRow struct {
	CtfName        string    `json:"ctf_name"`
	CtfDescription string    `json:"ctf_description"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	CtfAuthorID    int32     `json:"ctf_author_id"`
	Phrase         string    `json:"phrase"`
}

func (q *Queries) ListUsersJoinedCTFs(ctx context.Context, userID int32) ([]ListUsersJoinedCTFsRow, error) {
	rows, err := q.db.QueryContext(ctx, listUsersJoinedCTFs, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersJoinedCTFsRow
	for rows.Next() {
		var i ListUsersJoinedCTFsRow
		if err := rows.Scan(
			&i.CtfName,
			&i.CtfDescription,
			&i.StartDate,
			&i.EndDate,
			&i.CtfAuthorID,
			&i.Phrase,
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

const searchCTFs = `-- name: SearchCTFs :many
SELECT
    id, phrase, name, description, start_date, end_date, author_id
FROM
    ctfs
WHERE
    (? IS NULL OR ctfs.name LIKE CONCAT('%', ?, '%')) 
    and (? IS NULL OR ctfs.description LIKE CONCAT('%', ?, '%'))
ORDER BY start_date
`

type SearchCTFsParams struct {
	Name        interface{} `json:"name"`
	Description interface{} `json:"description"`
}

func (q *Queries) SearchCTFs(ctx context.Context, arg SearchCTFsParams) ([]Ctf, error) {
	rows, err := q.db.QueryContext(ctx, searchCTFs,
		arg.Name,
		arg.Name,
		arg.Description,
		arg.Description,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Ctf
	for rows.Next() {
		var i Ctf
		if err := rows.Scan(
			&i.ID,
			&i.Phrase,
			&i.Name,
			&i.Description,
			&i.StartDate,
			&i.EndDate,
			&i.AuthorID,
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

const updateCTF = `-- name: UpdateCTF :execresult
UPDATE ctfs
SET
    name = ?,
    description = ?,
    start_date = ?,
    end_date = ?
WHERE phrase = ?
`

type UpdateCTFParams struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Phrase      string    `json:"phrase"`
}

func (q *Queries) UpdateCTF(ctx context.Context, arg UpdateCTFParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateCTF,
		arg.Name,
		arg.Description,
		arg.StartDate,
		arg.EndDate,
		arg.Phrase,
	)
}
