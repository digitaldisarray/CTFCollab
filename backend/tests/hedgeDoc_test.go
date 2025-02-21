package tests

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"testing"
)

func TestCreateHedgeDocNote(t *testing.T) {
	noteID, err := createHedgeDocNote()
	if err != nil {
		t.Fatalf("Failed to create HedgeDoc note: %v", err)
	}

	fmt.Println("HedgeDoc Note Created! Note ID:", noteID)
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
	location := resp.Header.Get("Location")
	if location == "" {
		return "", fmt.Errorf("HedgeDoc API did not return a valid Location header")
	}

	// location might look like: "http://localhost:3001/zVl9cPTpStGxb5R05x3M5w"
	// Parse it to extract the note ID:
	u, err := url.Parse(location)
	if err != nil {
		return "", fmt.Errorf("failed to parse Location URL: %v", err)
	}

	// The note ID is basically everything after "/". EX:
	//   /zVl9cPTpStGxb5R05x3M5w   => zVl9cPTpStGxb5R05x3M5w
	noteID := path.Base(u.Path)
	if noteID == "" {
		return "", fmt.Errorf("could not extract note ID from location: %s", location)
	}

	return noteID, nil
}
