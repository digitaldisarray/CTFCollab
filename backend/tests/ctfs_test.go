package tests

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var (
	ctfs       []string // CTFs we have made to delete later
	adminToken string   // Admin token we can lazily use to nuke all ctfs at the end
)

func teardownCTFs() {
	for _, ctf := range ctfs {
		DeleteCTF(ctf, adminToken, client)
	}
}

// TODO: Split this up...
// Create a CTF and modify it in many ways before deleting it
func TestCTFEndToEnd(t *testing.T) {
	setup()
	defer teardownCTFs()
	defer teardownUsers()

	// ==================== CTF CREATION ====================
	// Login to admin account
	err := godotenv.Load("../../.env") // Load env
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	admin_username := os.Getenv("ADMIN_USERNAME") // Read admin creds
	admin_password := os.Getenv("ADMIN_PASSWORD")
	if admin_username == "" || admin_password == "" {
		t.Fatalf("Admin username or password not set in .env file")
	}

	admin_token, err := LoginToUser(admin_username, admin_password, client) // Sign in as admin
	if err != nil {
		t.Fatalf("Failed to login to admin user: %v", err)
	}

	// Create CTF
	t.Log("Creating CTF")
	phrase, err := CreateCTF(admin_token, RandomString(8), RandomString(20), client)
	if err != nil {
		t.Fatal(err)
	}
	ctfs = append(ctfs, phrase) // Add CTF to list to be deleted it later

	// Get the CTF to make sure it exists
	t.Log("Make sure CTF was created")
	body, err := GetCTF(admin_token, phrase, client)
	if err != nil {
		t.Fatal(err)
	}

	// ==================== JOIN CTF AS GUEST AND DO THINGS ====================
	t.Log("Joining CTF as a guest")
	guest_token, err := JoinCTFAsGuest(phrase, "TestGuest", client)
	if err != nil {
		t.Fatal(err)
	}

	// Make a challenge
	t.Log("Creating a challenge as a guest")
	challengeName := "Guest's Challenge"
	challengeDescription := "A description from guest"
	challengeFlag := "flag{guest}"
	createdName, hedgeDocURL, err := CreateChallenge(guest_token, phrase, challengeName, challengeDescription, challengeFlag, client)
	if err != nil {
		t.Fatalf("Failed to create challenge: %v", err)
	}
	t.Logf("Challenge created successfully! Name: %s, HedgeDoc URL: %s", createdName, hedgeDocURL)

	// Get challenges
	t.Log("Fetching challenges as a guest")
	challenges, err := GetCTFChallenges(guest_token, phrase, client)
	if err != nil {
		t.Fatalf("Failed to get challenges: %v", err)
	}
	if len(challenges) == 0 {
		t.Fatalf("Expected at least 1 challenge, got 0")
	}

	t.Log("Joining CTF as a guest with no nickname")
	_, err = JoinCTFAsGuest(phrase, "", client)
	if err != nil {
		t.Fatal(err)
	}

	// ==================== JOIN CTF AS NORMAL USER ====================
	// Create normal user
	t.Log("Creating user")
	username, password, err := CreateRandomUser(client)
	if err != nil {
		t.Fatal(err)
	}
	accounts = append(accounts, struct {
		username string
		password string
	}{username, password}) // Add accounts to list to delete later

	// Login to user
	t.Log("Logging into user")
	token, err := LoginToUser(username, password, client)
	if err != nil {
		t.Fatal(err)
	}

	// Join CTF
	if err := JoinCTF(token, phrase, client); err != nil {
		t.Fatal(err)
	}

	// ==================== CHALLENGE TESTING ====================
	// Create a challenge under the newly-created CTF
	challengeName = "Sample Challenge"
	challengeDescription = "A sample challenge description"
	challengeFlag = "flag{test}"
	createdName, hedgeDocURL, err = CreateChallenge(token, phrase, challengeName, challengeDescription, challengeFlag, client)
	if err != nil {
		t.Fatalf("Failed to create challenge: %v", err)
	}

	t.Logf("Challenge created successfully! Name: %s, HedgeDoc URL: %s", createdName, hedgeDocURL)

	// Create the second challenge
	secondName := "Another Challenge"
	secondDesc := "A second challenge description"
	secondFlag := "flag{test2}"
	createdName2, hedgeDocURL2, err := CreateChallenge(token, phrase, secondName, secondDesc, secondFlag, client)
	if err != nil {
		t.Fatalf("Failed to create second challenge: %v", err)
	}
	t.Logf("Second challenge created successfully! Name: %s, HedgeDoc URL: %s", createdName2, hedgeDocURL2)

	// Get challenges for the phrase and verify our new challenge is present
	t.Log("Retrieving challenges for CTF")
	challenges, err = GetCTFChallenges(token, phrase, client)
	if err != nil {
		t.Fatalf("Failed to get challenges: %v", err)
	}
	if len(challenges) == 0 {
		t.Fatalf("Expected at least 1 challenge, got 0")
	}

	t.Log("Listing all challenges:")
	for i, ch := range challenges {
		t.Logf("Challenge #%d: %+v", i+1, ch)
	}

	// Get challenge with id param and output it
	t.Log("Retrieving challenge with id")

	id := challenges[0].ChallengeID
	t.Log(token)
	challenge, err := GetChallenge(token, phrase, id, client)
	if err != nil {
		t.Fatalf("Failed to get challenge: %v", err)
	}

	t.Logf("Challenge id#%d: %+v", id, challenge)

	// ==================== CTF RENAME ====================
	// Rename CTF
	t.Log("Renaming CTF")
	new_name := RandomString(9)
	body["name"] = new_name
	err = UpdateCTF(token, phrase, body, client)
	if err != nil {
		t.Fatal(err)
	}

	// Get CTF again to see if name change took place
	body, err = GetCTF(token, phrase, client)
	if err != nil {
		t.Fatal(err)
	}
	if body["name"] != new_name {
		t.Fatal("New name not spotted after rename")
	}

	// Test Get all CTFs
	// List joined CTFs
	// Get CTF info

}
