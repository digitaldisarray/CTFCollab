package handler

import (
	"database/sql"
	"time"

	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	_ "github.com/go-sql-driver/mysql"
)

type Handler struct {
	Queries *db.Queries
}

func NewHandler(q *db.Queries) *Handler {
	return &Handler{Queries: q}
}

func LoadHandler(dbUrl string) (*Handler, error) {
	// Get a database handle
	conn, err := sql.Open("mysql", dbUrl)
	if err != nil {
		return nil, err
	}

	// See "Important settings" section.
	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	// Check handle to db
	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	// Make a handler object
	queries := db.New(conn)
	handler := NewHandler(queries)

	return handler, nil
}
