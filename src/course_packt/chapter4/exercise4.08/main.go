package main

import (
	"fmt"
	"os"
)

func getPassedArgs(minArgs int) []string {
	// Check if minimum arguments were passed
	// os.Args => Args hold the command-line arguments, starting with the program name.
	if len(os.Args) < minArgs {
		fmt.Printf("At least %v arguments are needed\n", minArgs)
		os.Exit(1)
	}

	// Declaring slice
	var args []string

	// Add all arguments in slice
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	// Returning slice
	return args
}

func findLongest(args []string) string {
	// Declaring string
	var longest string

	// Loop over the slice for find longest string
	for i := 0; i < len(args); i++ {
		// Find longest string
		if len(args[i]) > len(longest) {
			longest = args[i]
		}
	}

	// Returning longest string
	return longest
}

func main() {
	// Check if longest word is major that zero and passed major that three elements
	if longest := findLongest(getPassedArgs(3)); len(longest) > 0 {
		fmt.Println("The longest word passed was:", longest)
	} else {
		fmt.Println("There was an error")
		os.Exit(1)
	}
}
