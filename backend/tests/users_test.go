package tests

import (
	"fmt"
	"net/http"
	"testing"
)

var (
	accounts []struct {
		username string
		password string
	}
)

// Teardown function to clean up after tests
func teardownUsers() {
	for _, account := range accounts {
		// Login to user to get their token
		token, err := LoginToUser(account.username, account.password, client)
		if err != nil {
			fmt.Printf("Teardown: failed to login to user %s: %v\n", account.username, err)
			continue
		}

		// Delete the user
		DeleteUser(account.username, token, client)
	}
}

// TestPostUser creates a new user and also attempts to create a duplicate user
func TestPostUser(t *testing.T) {
	setup()
	defer teardownUsers()

	// Create a new user
	t.Log("Creating random user")
	username, password, err := CreateRandomUser(client)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("Created user successfully")
	}
	accounts = append(accounts, struct {
		username string
		password string
	}{username, password}) // Add accounts to list to delete later

	// Make sure we can't create the same user again
	t.Logf("Attempting to create user with same username: %s", username)
	if err := CreateUser(username, password, client); err == nil {
		t.Fatal("Created user with duplicate username")
	}

	t.Log("Successfully prevented duplicate user creation")
}

// TestLoginUser tests user login
func TestLoginUser(t *testing.T) {
	setup()
	defer teardownUsers()

	// Create a new user
	t.Log("Creating random user")
	username, password, err := CreateRandomUser(client)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("Created user successfully")
	}
	accounts = append(accounts, struct {
		username string
		password string
	}{username, password}) // Add accounts to list to delete later

	// Request to login
	t.Logf("Logging in with username: %s, password: %s", username, password)
	token, err := LoginToUser(username, password, client)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Received token: %s", token)
}

func TestChangePassword(t *testing.T) {
	setup()
	defer teardownUsers()

	// Create a new user
	t.Log("Creating random user")
	username, password, err := CreateRandomUser(client)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Created user successfully")

	// Request to login
	t.Logf("Logging in with username: %s, password: %s", username, password)
	token, err := LoginToUser(username, password, client)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Received token: %s", token)

	// Change the password
	t.Logf("Attempting to change password")
	newPassword := RandomString(16)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		SetBody(fmt.Sprintf(`{ "old_password": "%s", "new_password": "%s" }`, password, newPassword)).
		Post(fmt.Sprintf("http://localhost:1337/users/%s/password", username))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected status code %d, got %d. Response: %s", http.StatusOK, resp.StatusCode(), string(resp.Body()))
	}
	accounts = append(accounts, struct {
		username string
		password string
	}{username, newPassword}) // Add accounts to list to delete later

	// Verify new password by logging in
	t.Logf("Logging in with new password: %s", newPassword)
	token, err = LoginToUser(username, newPassword, client)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("New password worked, received token: %s", token)
}

// TestDeleteUser creates a new user and then deletes it
func TestDeleteUser(t *testing.T) {
	setup()

	// Create a new user
	t.Log("Creating random user")
	username, password, err := CreateRandomUser(client)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("Created user successfully")
	}

	// Log into user to get their token
	t.Logf("Logging in to user with username: %s", username)
	token, err := LoginToUser(username, password, client)
	if err != nil {
		t.Fatalf("Failed to login to user: %v", err)
	}

	// Delete the user
	t.Logf("Deleting user with username: %s", username)
	if err := DeleteUser(username, token, client); err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}

	t.Log("Deleted user successfully")
}

// TODO: Attempt to change password of another user
// TODO: Attempt to delete another user
// TODO: Attempt to login with incorrect password
