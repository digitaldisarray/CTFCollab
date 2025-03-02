package tests

import (
	"os"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

// Tries to login to the admin user specified in the .env file
func TestAdminUser(t *testing.T) {
	// Load the .env file located in the parent directory
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	// Read the 'ADMIN_USERNAME' and 'ADMIN_PASSWORD' from the .env file
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")
	if username == "" || password == "" {
		t.Fatalf("Admin username or password not set in .env file")
	}

	// Create a new Resty client
	client := resty.New()

	// Attempt to login to the user
	token, err := LoginToUser(username, password, client)
	if err != nil {
		t.Fatalf("Failed to login to admin user: %v", err)
	}

	// Check if the token is not empty
	if token == "" {
		t.Fatalf("Received empty token after login")
	}

	t.Logf("Successfully logged in as admin user, token: %s", token)
}
