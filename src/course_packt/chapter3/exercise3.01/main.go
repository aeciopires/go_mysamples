package main

import (
	"fmt"
	"unicode"
)

func passwordValidate(passwd string) bool {
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	result := []rune(passwd)
	if len(result) < 8 {
		return false
	}

	for _, character := range result {
		if unicode.IsUpper(character) {
			hasUpper = true
		}

		if unicode.IsLower(character) {
			hasLower = true
		}

		if unicode.IsNumber(character) {
			hasNumber = true
		}

		if unicode.IsPunct(character) || unicode.IsSymbol(character) {
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	if passwordValidate("") {
		fmt.Println("Password good")
	} else {
		fmt.Println("Password bad")
	}

	if passwordValidate("This!I5A") {
		fmt.Println("password good")
	} else {
		fmt.Println("password bad")
	}
}
