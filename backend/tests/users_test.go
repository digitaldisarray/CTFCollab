package tests

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
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
		token, err := loginToUser(account.username, account.password)
		if err != nil {
			fmt.Printf("Teardown: failed to login to user %s: %v\n", account.username, err)
			continue
		}

		// Delete the user
		deleteUser(account.username, token)
	}
}

// HELPER function to generate a random string
func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

// HELPER function to delete a user
func deleteUser(username, token string) error {
	// Send a DELETE request to the server
	url := fmt.Sprintf("http://localhost:1337/users/%s", username)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		Delete(url)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("Failed to delete user, status code: %d, Body: %s", resp.StatusCode(), string(resp.Body()))
	}

	return nil
}

// HELPER function to create a user with a random username and password
func createRandomUser() (string, string, error) {
	// Generate random username and password
	username := "testuser_" + randomString(8)
	password := randomString(16)

	// Crete the user
	if err := createUser(username, password); err != nil {
		return "", "", err
	}

	return username, password, nil
}

// HELPER function to create a user with a given username and password
func createUser(username, password string) error {
	// Send a POST request to the server
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password)).
		Post("http://localhost:1337/user")

	// Check for errors
	if err != nil {
		return fmt.Errorf("Failed to create user: %v\n", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("Failed to create user, status code: %d\n", resp.StatusCode())
	}

	return nil
}

func loginToUser(username, password string) (string, error) {
	// Request to login
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password)).
		Post("http://localhost:1337/user/login")

	// Check for errors
	if err != nil {
		return "", fmt.Errorf("Failed to send request: %v\n", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("Failed to login, status code: %d, body: %s\n", resp.StatusCode(), string(resp.Body()))
	}

	// Check that the response body contains a JWT token
	var responseBody map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &responseBody); err != nil {
		return "", fmt.Errorf("Failed to parse response body: %v\n", err)
	}
	token, exists := responseBody["token"]
	if !exists {
		return "", fmt.Errorf("Expected to find 'token' in response body, but it was missing.")
	}

	return token.(string), nil
}

// TestPostUser creates a new user and also attempts to create a duplicate user
func TestPostUser(t *testing.T) {
	setup()
	defer teardownUsers()

	// Create a new user
	t.Log("Creating random user")
	username, password, err := createRandomUser()
	if err != nil {
		t.Fatal("Failed to create user")
	} else {
		t.Log("Created user successfully")
	}
	accounts = append(accounts, struct {
		username string
		password string
	}{username, password}) // Add accounts to list to delete later

	// Make sure we can't create the same user again
	t.Logf("Attempting to create user with same username: %s", username)
	if err := createUser(username, password); err == nil {
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
	username, password, err := createRandomUser()
	if err != nil {
		t.Fatal("Failed to create user")
	} else {
		t.Log("Created user successfully")
	}
	accounts = append(accounts, struct {
		username string
		password string
	}{username, password}) // Add accounts to list to delete later

	// Request to login
	t.Logf("Logging in with username: %s, password: %s", username, password)
	token, err := loginToUser(username, password)
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
	username, password, err := createRandomUser()
	if err != nil {
		t.Fatal("Failed to create user")
	}
	t.Log("Created user successfully")

	// Request to login
	t.Logf("Logging in with username: %s, password: %s", username, password)
	token, err := loginToUser(username, password)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Received token: %s", token)

	// Change the password
	newPassword := randomString(16)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		SetBody(fmt.Sprintf(`{
			"username": "%s",
			"password": "%s"
		}`, username, newPassword)).
		Post("http://localhost:1337/users/password")
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
	token, err = loginToUser(username, newPassword)
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
	username, password, err := createRandomUser()
	if err != nil {
		t.Fatal("Failed to create user")
	} else {
		t.Log("Created user successfully")
	}

	// Log into user to get their token
	t.Logf("Logging in to user with username: %s", username)
	token, err := loginToUser(username, password)
	if err != nil {
		t.Fatalf("Failed to login to user: %v", err)
	}

	// Delete the user
	t.Logf("Deleting user with username: %s", username)
	if err := deleteUser(username, token); err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}

	t.Log("Deleted user successfully")
}

// TODO: Attempt to change password of another user
// TODO: Attempt to delete another user
// TODO: Attempt to login with incorrect password
