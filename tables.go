package main

import (
	"fmt"

	db "github.com/aerospike/aerospike-client-go"
)

/**
 * Clear a database set of all records
 *
 * The function can only do this if a correct WritePolicy was issued when the records were added to the set.
 * Specifically, the SendKey value of the policy must be True or the server will not return a record's primary
 * key, which is used to delete the records.
 */
func clearTable(table string) (int, error) {
	c := 0

	// Query the server for all record of corresponding set
	stmt := db.NewStatement(ARG_DB_NAMESPACE, table)
	rs, err := DBConn.Query(nil, stmt)

	if err != nil {
		return c, err
	}

	fmt.Printf("\nReading set => %v:\n", table)

	// Iterate through records
	for res := range rs.Results() {
		record := res.Record.Key.Value().String()

		// Check if server returns the primary key
		if record == "" {
			fmt.Printf("\tRecord %v: ... aborted (no primary key found)\n", c)
			continue
		}

		// Delete the record from the server
		fmt.Printf("\tRecord %v: ... deleting => %v\n", c+1, record)
		DBConn.Delete(nil, res.Record.Key)

		// Increase record count
		c++
	}

	if c == 0 {
		fmt.Println("\tNo records found...\n")
	}

	return c, nil
}
