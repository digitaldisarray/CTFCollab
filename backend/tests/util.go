package tests

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// Helper function to generate a random string
func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

// HELPER function to delete a user
func DeleteUser(username, token string, client *resty.Client) error {
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
		return fmt.Errorf("failed to delete user, status code: %d, body: %s", resp.StatusCode(), string(resp.Body()))
	}

	return nil
}

// HELPER function to create a user with a random username and password
func CreateRandomUser(client *resty.Client) (string, string, error) {
	// Generate random username and password
	username := "testuser_" + RandomString(8)
	password := RandomString(16)

	// Crete the user
	if err := CreateUser(username, password, client); err != nil {
		return "", "", err
	}

	return username, password, nil
}

// HELPER function to create a user with a given username and password
func CreateUser(username, password string, client *resty.Client) error {
	// Send a POST request to the server
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password)).
		Post("http://localhost:1337/users")

	// Check for errors
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("failed to create user, status code: %d, body: %s", resp.StatusCode(), string(resp.Body()))
	}

	return nil
}

func LoginToUser(username, password string, client *resty.Client) (string, error) {
	// Request to login
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password)).
		Post("http://localhost:1337/users/login")

	// Check for errors
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("failed to login, status code: %d, body: %s", resp.StatusCode(), string(resp.Body()))
	}

	// Check that the response body contains a JWT token
	var responseBody map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &responseBody); err != nil {
		return "", fmt.Errorf("failed to parse response body: %v", err)
	}
	token, exists := responseBody["token"]
	if !exists {
		return "", fmt.Errorf("expected to find 'token' in response body, but it was missing")
	}

	return token.(string), nil
}

func DeleteCTF(phrase, token string, client *resty.Client) error {
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		Delete(fmt.Sprintf("http://localhost:1337/ctfs/%s", phrase))

	// Check for errors
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("failed to delete ctf, status code: %d, body: %s", resp.StatusCode(), string(resp.Body()))
	}

	return nil
}

func CreateCTF(token, name, description string, client *resty.Client) (string, error) {
	start_date := "2025-01-31T09:00:00Z"
	end_date := "2025-02-01T12:00:00Z"

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		SetBody(fmt.Sprintf(`{
			"name": "%s",
			"description": "%s",
			"start_date": "%s",
			"end_date": "%s"
		}`, name, description, start_date, end_date)).
		Post("http://localhost:1337/ctfs")

	// Check for errors
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("failed to create ctf, status code: %d, body: %s", resp.StatusCode(), string(resp.Body()))
	}

	// Parse response body
	var responseBody map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &responseBody); err != nil {
		return "", fmt.Errorf("failed to parse response body: %v", err)
	}
	phrase, exists := responseBody["phrase"]
	if !exists {
		return "", fmt.Errorf("expected to find 'phrase' in response body, but it was missing")
	}

	return phrase.(string), nil
}

func UpdateCTF(token, phrase string, new_body map[string]interface{}, client *resty.Client) error {
	// Convert map to JSON string
	jsonBytes, err := json.Marshal(new_body)
	if err != nil {
		return fmt.Errorf("error marshalling new body map to JSON: %v", err)
	}

	// Put request to localhost:1337/ctfs/<phrase>
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		SetBody(string(jsonBytes)).
		Put(fmt.Sprintf("http://localhost:1337/ctfs/%s", phrase))

	// Check for errors
	if err != nil {
		return fmt.Errorf("failed to send rename ctf request: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("failed to rename ctf, status code: %d, body: %s", resp.StatusCode(), string(resp.Body()))
	}

	return nil
}

func GetCTF(token, phrase string, client *resty.Client) (map[string]interface{}, error) {
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		Get(fmt.Sprintf("http://localhost:1337/ctfs/%s", phrase))

	// Check for errors
	if err != nil {
		return nil, fmt.Errorf("failed to get ctf: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to get ctf, status code: %d, body: %s", resp.StatusCode(), string(resp.Body()))
	}

	// Parse and return response body
	var responseBody map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &responseBody); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %v", err)
	}

	return responseBody, nil
}

func JoinCTF(token, phrase string, client *resty.Client) error {
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		Post(fmt.Sprintf("http://localhost:1337/ctfs/%s/join", phrase))

	// Check for errors
	if err != nil {
		return fmt.Errorf("failed to join ctf: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("failed to join ctf, status code: %d, body: %s", resp.StatusCode(), string(resp.Body()))
	}

	return nil
}

// JoinCTFAsGuest sends a request to join a CTF as a guest.
// It takes the CTF phrase, a nickname, and a Resty client as input.
// Returns the token as a string or an error if the request fails.
func JoinCTFAsGuest(phrase, nickname string, client *resty.Client) (string, error) {
	// Define the request payload
	payload := map[string]string{
		"nickname": nickname,
	}

	if nickname == "" {
		payload = nil
	}

	// Make the POST request to the endpoint
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(fmt.Sprintf("http://localhost:1337/ctfs/%s/join-as-guest", phrase))
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}

	// Check if the response status code is not 200 OK
	if resp.StatusCode() != 200 {
		return "", fmt.Errorf("failed to join CTF as guest: %s", resp.String())
	}

	// Extract the token from the response cookies
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "token" {
			return cookie.Value, nil
		}
	}

	return "", fmt.Errorf("token not found in response cookies")
}

type ChallengeResponse struct {
	ChallengeName string `json:"challenge_name"`
	HedgeDocURL   string `json:"hedgedoc_url"`
}

// CreateChallenge creates a challenge via POST /:phrase/challenges using resty.Client.
// It returns the challenge name and HedgeDoc URL if successful.
func CreateChallenge(token, phrase, name, description, flag string, client *resty.Client) (string, string, error) {
	// Build the request body
	body := map[string]string{
		"name":        name,
		"description": description,
		"flag":        flag,
	}

	// Perform the POST request
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(fmt.Sprintf("http://localhost:1337/ctfs/%s/challenges", phrase))
	if err != nil {
		return "", "", fmt.Errorf("failed to create challenge: %w", err)
	}

	// Check the response status
	if resp.StatusCode() != http.StatusOK {
		return "", "", fmt.Errorf(
			"unexpected status code %d when creating challenge: %s",
			resp.StatusCode(),
			resp.String(),
		)
	}

	// Parse the response body into ChallengeResponse
	var challengeResp ChallengeResponse
	if err := json.Unmarshal(resp.Body(), &challengeResp); err != nil {
		return "", "", fmt.Errorf("failed to unmarshal challenge response: %w", err)
	}

	return challengeResp.ChallengeName, challengeResp.HedgeDocURL, nil
}

// GetChallengesResponse matches the JSON structure returned by the API endpoint GET /ctfs/:phrase/challenges.
type GetCTFChallengesRow struct {
	ChallengeID          int32        `json:"challenge_id"`
	ChallengeName        string       `json:"challenge_name"`
	ChallengeDescription string       `json:"challenge_description"`
	Flag                 string       `json:"flag"`
	ChallengeCreatedAt   sql.NullTime `json:"challenge_created_at"`
	HedgedocURL          string       `json:"hedgedoc_url"`
}

// GetCTFChallenges makes a GET request to /ctfs/:phrase/challenges and returns the list of challenges.
func GetCTFChallenges(token, phrase string, client *resty.Client) ([]GetCTFChallengesRow, error) {
	url := fmt.Sprintf("http://localhost:1337/ctfs/%s/challenges", phrase)

	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to send GET request for challenges: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf(
			"unexpected status code %d from GetCTFChallenges, body: %s",
			resp.StatusCode(),
			resp.String(),
		)
	}

	var challenges []GetCTFChallengesRow
	if err := json.Unmarshal(resp.Body(), &challenges); err != nil {
		return nil, fmt.Errorf("failed to unmarshal get challenges response: %w", err)
	}

	return challenges, nil
}

func GetChallenge(token string, phrase string, id int32, client *resty.Client) (map[string]interface{}, error) {
	url := fmt.Sprintf("http://localhost:1337/ctfs/%s/challenge/%d", phrase, id)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to send GET request for challenge: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf(
			"unexpected status code %d from GetChallenge, body: %s",
			resp.StatusCode(),
			resp.String(),
		)
	}
	var responseBody map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &responseBody); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %v", err)
	}

	return responseBody, nil
}
