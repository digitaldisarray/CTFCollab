package handler

import (
	"database/sql"
	"log"
	"time"

	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Import the file source driver
)

type Handler struct {
	Queries *db.Queries
}

func NewHandler(q *db.Queries) *Handler {
	return &Handler{Queries: q}
}

func LoadHandler(dbUrl string) (*Handler, error) {
	// Get a database handle
	handle, err := sql.Open("mysql", dbUrl)
	if err != nil {
		return nil, err
	}
	handle.SetConnMaxLifetime(time.Minute * 3)
	handle.SetMaxOpenConns(10) // Subject to change
	handle.SetMaxIdleConns(10)

	// Check handle to db
	if err = handle.Ping(); err != nil {
		return nil, err
	}

	// Do database migrations
	driver, err := mysql.WithInstance(handle, &mysql.Config{})
	if err != nil {
		log.Fatalf("Error creating MySQL driver: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"ctfcollab",
		driver,
	)
	if err != nil {
		log.Fatalf("Error creating migration instance: %v", err)
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error spinning up schema: %v", err)
	}

	// Make a handler object
	queries := db.New(handle)
	handler := NewHandler(queries)

	return handler, nil
}
