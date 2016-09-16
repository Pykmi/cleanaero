package main

import (
	"fmt"

	db "github.com/aerospike/aerospike-client-go"
)

func clearTable(table string) error {
	stmt := db.NewStatement(ARG_DB_NAMESPACE, table)
	rs, _ := DBConn.Query(nil, stmt)

	c := 1
	var key string

	fmt.Printf("\nReading set => %v:\n", table)

	for res := range rs.Results() {
		key = res.Record.Key.String()

		if key == "" {
			fmt.Printf("\tRecord %v: ... aborted (no key found)\n", c)
			continue
		}

		fmt.Printf("\tRecord %v: ... aborted (no key found)\n", c)
		c++
	}

	return nil
}
