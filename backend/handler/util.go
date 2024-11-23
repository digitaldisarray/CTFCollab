package handler

import (
	"database/sql"
	"fmt"
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
