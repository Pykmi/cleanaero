package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	db "github.com/aerospike/aerospike-client-go"
)

/**
 * Globals for the command-line arguments.
 */
var ARG_DB_PORT int
var ARG_DB_HOST string
var ARG_DB_NAMESPACE string
var ARG_DB_SETS string
var ARG_VERBOSE bool

/**
 * Global database connection
 */
var DBConn *db.Client

func main() {
	count := 0

	// Initialize command-line arguments
	Init()

	// Open a new persistent database connection
	C, err := db.NewClient(ARG_DB_HOST, ARG_DB_PORT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Save database connection to global variable
	DBConn = C
	defer DBConn.Close()

	// Run verbose for main()
	initV()

	// Parse set names from argument
	sets := parseTableNames()

	// Check for errors in set names
	if sets == nil {
		fmt.Fprintf(os.Stderr, "\nError: no sets specified.\n")
		os.Exit(1)
	}

	for i, _ := range sets {
		n, err := clearTable(sets[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			continue
		}

		count = count + n
	}

	fmt.Printf("\nComplete... removed %v records", count)
}

/**
 * Initialize command-line arguments and parse them.
 */
func Init() {
	// Set parameters for database port and hostname
	flag.IntVar(&ARG_DB_PORT, "port", 3000, "Database server port number.")
	flag.StringVar(&ARG_DB_HOST, "host", "localhost", "Database server hostname.")

	// Set parameters for the database namespace and set names
	flag.StringVar(&ARG_DB_NAMESPACE, "ns", "test", "Namespace to access within the database.")
	flag.StringVar(&ARG_DB_SETS, "sets", "", "Sets to wipe clean in the database. Separate multiple sets by colon.")

	// Additional parameter for verbose program output
	flag.BoolVar(&ARG_VERBOSE, "v", false, "Make program output verbose.")

	// Parse the arguments
	flag.Parse()
}

/**
 * Parse the set names from the passed argument.
 *
 * Each set or table name is separated by a colon. If no colon is found but argument has a value, it is presumed to be
 * the name of one set or table name.
 */
func parseTableNames() []string {
	if strings.Contains(ARG_DB_SETS, ",") == false {
		if len(ARG_DB_SETS) > 0 {
			ss := []string{ARG_DB_SETS}
			parseTableNamesV(ss)
			return ss
		}

		return nil
	}

	s := strings.Split(ARG_DB_SETS, ",")
	parseTableNamesV(s)
	return s
}

/**
 * Panic for errors. Should only be used in non-critical times, like when server is started.
 */
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
