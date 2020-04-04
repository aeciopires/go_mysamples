package main

import (
	"fmt"
	"os"
)

func getPassedArgs() []string {
	var args []string

	// os.Args => Args hold the command-line arguments, starting with the program name.
	// Add arguments of user in slice
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	// Returning slice
	return args
}

func getLocals(extraLocals []string) []string {
	// Create local variable
	var locales []string

	// Add two elements is slice
	locales = append(locales, "en_US", "fr_FR")

	// Add arguments of user in slice
	locales = append(locales, extraLocals...)

	// Returning slice
	return locales
}

func main() {
	// Add arguments of user and elements in slice of main function
	locales := getLocals(getPassedArgs())
	fmt.Println("Locales to use:", locales)
}
