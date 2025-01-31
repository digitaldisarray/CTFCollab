package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type CtfPostResponse struct {
	CtfID int `json:"ctf_id"`
}

type CtfGetResponse struct {
	Name        string `json:"name"`
	Phrase      string `json:"phrase"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	AuthorID    int32  `json:"author_id"`
}

var (
	ctf_id      int
	name        = "Testing"
	phrase      = "testphrase"
	description = "This is a test"
	start_date  = "2025-01-31T09:00:00Z"
	end_date    = "2025-02-01T12:00:00Z"
	author_id   = 3
)

func teardownCTF() {
	if user_id != 0 {
		deleteCTF(phrase)
	}
}

func getCtf(ctfPhrase string) CtfGetResponse {
	url := fmt.Sprintf("http://localhost:1337/ctfs/%s", ctfPhrase)

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get(url)

	if err != nil {
		fmt.Printf("Failed to send request: %s", err)
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Printf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode())
	}

	var ctfResponse CtfGetResponse
	err = json.Unmarshal(resp.Body(), &ctfResponse)
	if err != nil {
		fmt.Printf("Failed to unmarshal response: %sv", err)
	}

	return ctfResponse

}
func deleteCTF(ctfPhrase string) {
	url := fmt.Sprintf("http://localhost:1337/ctfs/%s", ctfPhrase)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Delete(url)

	if err != nil {
		fmt.Printf("Failed to delete CTF: %v\n", err)
		return
	}

	if resp.StatusCode() != http.StatusOK {
		fmt.Printf("Failed to delete CTF, status code: %d\n", resp.StatusCode())
	}
}
func TestPostCtf(t *testing.T) {
	setup()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{
		"name": "%s",
		"phrase": "%s",
		"description": "%s",
		"start_date": "%s",
		"end_date": "%s",
		"author_id": %d}`, name, phrase, description, start_date, end_date, author_id)).
		Post("http://localhost:1337/ctfs")

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode())
	}

	var ctfResponse CtfPostResponse
	err = json.Unmarshal(resp.Body(), &ctfResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	ctf_id = ctfResponse.CtfID
	t.Logf("Created CTF with ID: %d", ctf_id)
}

func TestGetCtf(t *testing.T) {
	setup()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get("http://localhost:1337/ctfs/testphrase")

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode())
	}

	var ctfResponse CtfGetResponse
	err = json.Unmarshal(resp.Body(), &ctfResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if !(ctfResponse.Name == name &&
		ctfResponse.Description == description &&
		ctfResponse.AuthorID == int32(author_id) &&
		ctfResponse.StartDate == start_date &&
		ctfResponse.EndDate == end_date &&
		ctfResponse.Phrase == phrase) {
		t.Errorf("Incorrect response.")
	}

}

func TestUpdateCtf(t *testing.T) {
	setup()
	defer teardownCTF()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{
			"name": "%s",
			"phrase": "%s",
			"description": "Updated",
			"start_date": "%s",
			"end_date": "%s"}`, name, phrase, start_date, end_date)).
		Put(fmt.Sprintf("http://localhost:1337/ctfs/%s", phrase))

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode())
	}

	var updatedCtf = getCtf(phrase)
	if updatedCtf.Description != "Updated" {
		t.Errorf("Failed to update CTF")
	}
}

func TestDeleteCtf(t *testing.T) {
	setup()
	defer teardownCTF()
	if ctf_id == 0 {
		t.Skip("No CTF ID available to delete")
	}
	url := fmt.Sprintf("http://localhost:1337/ctfs/%s", phrase)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Delete(url)

	if err != nil {
		fmt.Printf("Failed to delete CTF: %v\n", err)
		return
	}

	if resp.StatusCode() != http.StatusOK {
		fmt.Printf("Failed to delete CTF, status code: %d\n", resp.StatusCode())
	}
}
