package tests

import "testing"

var (
	ctfs       []string // CTFs we have made to delete later
	adminToken string   // Admin token we can lazily use to nuke all ctfs at the end
)

func teardownCTFs() {
	for _, ctf := range ctfs {
		DeleteCTF(ctf, adminToken, client)
	}
}

// Create a CTF and modify it in many ways before deleting it
func TestCTFEndToEnd(t *testing.T) {
	setup()
	defer teardownCTFs()
	defer teardownUsers()

	// Create user
	t.Log("Creating user")
	username, password, err := CreateRandomUser(client)
	if err != nil {
		t.Fatal(err)
	}
	accounts = append(accounts, struct {
		username string
		password string
	}{username, password}) // Add accounts to list to delete later

	// Give user admin
	t.Log("Making user admin")
	err = MakeUserAdmin(username, client)
	if err != nil {
		t.Fatal(err)
	}

	// Login to user
	t.Log("Logging into user")
	token, err := LoginToUser(username, password, client)
	if err != nil {
		t.Fatal(err)
	}
	adminToken = token

	// Create CTF
	t.Log("Creating CTF")
	phrase, err := CreateCTF(token, RandomString(8), RandomString(20), client)
	if err != nil {
		t.Fatal(err)
	}
	ctfs = append(ctfs, phrase) // Add CTF to list to be deleted it later

	// Rename CTF
	// Test Get all CTFs

	// Create non privileged user
	// Have non privileged user join ctf
	// List joined CTFs
	// Get CTF info

}
