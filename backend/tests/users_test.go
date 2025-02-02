package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type UserResponse struct {
	UserID int `json:"user_id"`
}

var (
	user_id  int
	password = "xyz"
)

// Teardown function to clean up after tests
func teardownUsers() {
	if user_id != 0 {
		deleteUser(user_id)
	}
}

// Helper function to delete a user
func deleteUser(userID int) {
	url := fmt.Sprintf("http://localhost:1337/users/%d", userID)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Delete(url)

	if err != nil {
		fmt.Printf("Failed to delete user: %v\n", err)
		return
	}

	if resp.StatusCode() != http.StatusOK {
		fmt.Printf("Failed to delete user, status code: %d\n", resp.StatusCode())
	}
}

// TestPostUser creates a new user
func TestPostUser(t *testing.T) {
	setup()

	username := "testuser"

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"username": "%s", "password_hash": "xyz"}`, username)).
		Post("http://localhost:1337/user")

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode())
	}

	var userResponse UserResponse
	err = json.Unmarshal(resp.Body(), &userResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	user_id = userResponse.UserID
	t.Logf("Created user with ID: %d", user_id)
}

// TestLoginUser tests user login
func TestLoginUser(t *testing.T) {
	setup()
	defer teardownUsers()

	username := "testuser"

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"username": "%s", "password_hash": "%s"}`, username, password)).
		Post("http://localhost:1337/user/login")

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode())
	}
}

func TestChangePassword(t *testing.T) {
	setup()
	defer teardownUsers()

	// Change the password
	newPassword := "abc"

	t.Logf("Changing password for user ID: %d, new password: %s", user_id, newPassword)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{
			"id": %d,
			"password_hash": "%s"
		}`, user_id, newPassword)).
		Post("http://localhost:1337/users/password")

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	t.Logf("ChangePassword response: %s", resp.Body()) // Log the response

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected status code %d, got %d. Response: %s", http.StatusOK, resp.StatusCode(), resp.Body())
	}

	username := "testuser"

	// Verify the new password by logging in
	loginPayload := fmt.Sprintf(`{"username": "%s", "password_hash": "%s"}`, username, newPassword)
	t.Logf("Login request payload: %s", loginPayload) // Log the login payload

	resp, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(loginPayload).
		Post("http://localhost:1337/user/login")

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	t.Logf("Login response: %s", resp.Body()) // Log the login response

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected status code %d, got %d. Response: %s", http.StatusOK, resp.StatusCode(), resp.Body())
	}
}

// TestDeleteUser deletes the user created in TestPostUser
func TestDeleteUser(t *testing.T) {
	setup()
	defer teardownUsers()

	if user_id == 0 {
		t.Skip("No user ID available to delete")
	}

	url := fmt.Sprintf("http://localhost:1337/users/%d", user_id)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Delete(url)

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode())
	}

	user_id = 0 // Reset user_id after deletion
}
