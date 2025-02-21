package handler

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/alexedwards/argon2id"
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

	// Create an admin account if specified and we haven't already
	CreateAdmin(queries)

	return handler, nil
}

func CreateAdmin(queries *db.Queries) {
	admin_username, found := os.LookupEnv("ADMIN_USERNAME")
	if !found || admin_username == "" {
		log.Println("ADMIN_USERNAME environment variable not set, skipping creation")
		return
	}

	admin_password, found := os.LookupEnv("ADMIN_PASSWORD")
	if !found || admin_password == "" {
		log.Println("ADMIN_PASSWORD environment variable not set, skipping creation")
		return
	}

	// Make sure admin account doesn't already exist
	ctx := context.Background()
	_, err := queries.GetUserByUsername(ctx, admin_username)
	if err == nil {
		log.Printf("ADMIN_USERNAME: %s already exists, skipping creation\n", admin_username)
		return
	}

	// Hash admin password
	hash, err := argon2id.CreateHash(admin_password, argon2id.DefaultParams)
	if err != nil {
		log.Printf("Error hashing admin password: %v", err)
		return
	}

	// Create the admin account
	params := new(db.CreateAdminParams)
	params.Username = admin_username
	params.PasswordHash = hash
	queries.CreateAdmin(ctx, *params)
}
