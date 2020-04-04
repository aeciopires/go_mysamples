package main

import (
	"fmt"
	"os"
	"strings"
)

// Creating global variable of type struct
type locale struct {
	language  string
	territory string
}

// Function that returns the test data
func getLocales() map[locale]struct{} {
	supportedLocales := make(map[locale]struct{}, 5)
	supportedLocales[locale{"en", "US"}] = struct{}{}
	supportedLocales[locale{"en", "CN"}] = struct{}{}
	supportedLocales[locale{"fr", "CN"}] = struct{}{}
	supportedLocales[locale{"fr", "FR"}] = struct{}{}
	supportedLocales[locale{"ru", "RU"}] = struct{}{}
	return supportedLocales
}

// Create a function that uses a passed local struct
// to check the sample data to see if a locale exists
func localeExists(l locale) bool {
	_, exists := getLocales()[l]
	return exists
}

func main() {

	// Check that an argument has been passed
	if len(os.Args) < 2 {
		fmt.Println("No locale passed")
		os.Exit(1)
	}

	// Process the passed argument to make sure it's in a valid format
	localeParts := strings.Split(os.Args[1], "_")
	if len(localeParts) != 2 {
		fmt.Printf("Invalid locale passed: %v\n", os.Args[1])
		os.Exit(1)
	}

	// Create a local struct value using the passed argument data
	passedLocale := locale{
		territory: localeParts[1],
		language:  localeParts[0],
	}

	// Call the function and print an error message if it doesn't exist;
	// otherwise, print that the locale is supported
	if !localeExists(passedLocale) {
		fmt.Printf("Locale not supported: %v\n", os.Args[1])
		os.Exit(1)
	}

	fmt.Println("Locale passed is supported")
}
