package handler

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

// Makes sure the expected vs actual rows affected match in a sql query result
// Returns the last inserted ID. Returns -1 and an error in the event of a failure
func VerifyParseResult(result sql.Result, expected_affected int64) (int64, error) {
	// Get the number of rows affected
	rows_affected, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	// Make sure the expected vs actual rows affected match
	if rows_affected != expected_affected {
		return -1, fmt.Errorf("mismatch in expected vs actual rows affected (%d vs %d)", expected_affected, rows_affected)
	}

	// Get the last inserted id
	last_inserted, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return last_inserted, nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateGuestNickname generates a random nickname in the format "guest_<random_string>"
func GenerateGuestNickname() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomString := make([]byte, 6)
	for i := range randomString {
		randomString[i] = charset[rng.Intn(len(charset))]
	}
	return "guest_" + string(randomString)
}
