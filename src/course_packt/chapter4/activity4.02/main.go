package main

import (
	"fmt"
	"os"
)

// Create global variable with map of strings
var users = map[string]string{
	"305": "Sue",
	"204": "Bob",
	"631": "Jake",
	"073": "Tracy",
}

// Function that returns the user's name and whether it exists
func getName(id string) (string, bool) {
	name, exists := users[id]
	return name, exists
}

/*
 Check the passed arguments. Call the function, print if there's an
 error, and exit if the user doesn't exist. Print a greeting to the user if they do exist
*/
func main() {
	// os.Args => Args hold the command-line arguments, starting with the program name.
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}

	name, exists := getName(os.Args[1])

	if !exists {
		fmt.Printf("error: user (%v) not found", os.Args[1])
		os.Exit(1)
	}

	fmt.Println("Hi,", name)
}
