package handler

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/alexedwards/argon2id"
	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	"github.com/digitaldisarray/ctfcollab/websocket"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Import the file source driver
)

type Handler struct {
	Queries   *db.Queries
	JWTSecret string
	WsHub     *websocket.Hub
}

// NewHandler now correctly accepts the Queries and Hub
func NewHandler(q *db.Queries, wsHub *websocket.Hub) *Handler {
	return &Handler{
		Queries: q,
		WsHub:   wsHub,
		// JWTSecret will be set later in LoadHandler
	}
}

func LoadHandler(dbUrl, jwtSecret string) (*Handler, error) {
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
		handle.Close() // Close handle on ping error
		return nil, err
	}

	// Do database migrations
	driver, err := mysql.WithInstance(handle, &mysql.Config{})
	if err != nil {
		handle.Close()
		log.Printf("Error creating MySQL driver: %v", err)
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"ctfcollab",
		driver,
	)
	if err != nil {
		handle.Close()
		log.Printf("Error creating migration instance: %v", err)
		return nil, err
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		ver, dirty, dirtyErr := m.Version()
		if dirtyErr == nil && dirty {
			log.Printf("Database is dirty at version %d.", ver)
		} else {
			handle.Close()
			log.Printf("Error spinning up schema: %v", err)
			return nil, err
		}

	}
	log.Println("Database migrations applied successfully or no changes needed.")

	// Create and Start WebSocket Hub
	wsHub := websocket.NewHub()
	go wsHub.Run() // Start the hub's event loop in a background goroutine
	log.Println("WebSocket Hub created and started.")

	// Make a handler object, passing queries and the hub
	queries := db.New(handle)
	handler := NewHandler(queries, wsHub) // Passing wsHub here

	// Set the JWT secret
	handler.JWTSecret = jwtSecret

	// Create an admin account if specified and we haven't already
	CreateAdmin(queries)

	log.Println("Handler loaded successfully.")
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Add timeout
	defer cancel()

	_, err := queries.GetUserByUsername(ctx, admin_username)
	if err == nil {
		log.Printf("Admin user '%s' already exists, skipping creation.\n", admin_username)
		return
	}
	if err != sql.ErrNoRows {
		log.Printf("Error checking for existing admin user '%s': %v", admin_username, err)
		return
	}
	// If we reach here, err was sql.ErrNoRows, so the user doesn't exist - proceed.

	// Hash admin password
	hash, err := argon2id.CreateHash(admin_password, argon2id.DefaultParams)
	if err != nil {
		log.Printf("Error hashing admin password: %v", err)
		return
	}

	// Create the admin account
	params := db.CreateAdminParams{
		Username:     admin_username,
		PasswordHash: hash,
	}

	// Use a new context with timeout for the creation query
	createCtx, createCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer createCancel()

	_, err = queries.CreateAdmin(createCtx, params)
	if err != nil {
		log.Printf("Error creating admin user '%s': %v", admin_username, err)
		return
	}

	log.Printf("Admin user '%s' created successfully.", admin_username)
}
