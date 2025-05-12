package tests

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestCreateHedgeDocNote(t *testing.T) {
	hurl, err := createHedgeDocNote()
	if err != nil {
		t.Fatalf("Failed to create HedgeDoc note: %v", err)
	}

	t.Log("HedgeDoc Note Created! Note ID:", hurl)
}

func createHedgeDocNote() (string, error) {
	hedgedocAPI := "http://localhost:3001/new"

	req, err := http.NewRequest("POST", hedgedocAPI, strings.NewReader("# Challenge Notes"))
	if err != nil {
		return "", fmt.Errorf("failed to create HedgeDoc request: %v", err)
	}

	// Build a custom client that doesn't follow redirects.
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request to HedgeDoc: %v", err)
	}
	defer resp.Body.Close()

	// Should now have the response. Extract the 'Location' header:
	// location might look like: "http://localhost:3001/zVl9cPTpStGxb5R05x3M5w"
	location := resp.Header.Get("Location")
	if location == "" {
		return "", fmt.Errorf("HedgeDoc API did not return a valid Location header")
	}

	return location, nil
}
