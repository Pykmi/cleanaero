package main

import "fmt"

func initV() bool {
	if ARG_VERBOSE == false {
		return false
	}

	fmt.Println("Initializing setting...\n")
	fmt.Printf("Server\t\t\t => %v:%v\n", ARG_DB_HOST, ARG_DB_PORT)
	fmt.Printf("Namespace\t\t => %v\n", ARG_DB_NAMESPACE)

	return true
}

func parseTableNamesV(sets []string) bool {
	if ARG_VERBOSE == false {
		return false
	}

	fmt.Printf("Registered %v sets\t => ", len(sets))
	for i, _ := range sets {
		fmt.Printf("%v ", sets[i])
	}

	fmt.Println("")
	return true
}
