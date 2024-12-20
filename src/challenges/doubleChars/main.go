/*
Reference: https://tutorialedge.net/challenges/go/repeating-letters/
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DoubleChars(original string) string {
	newString := ""
	for _, char := range original {
		newString += string(char)
		newString += string(char)
	}

	return newString
}

func main() {
	fmt.Println("*********************")
	fmt.Println("Double Letters")
	fmt.Println("Show another string which has every letter in the word doubled.")
	fmt.Println("*********************")

	original := "gophers"
	doubled := DoubleChars(original)
	fmt.Println("Original:", original, "...", "Result: ", doubled) // ggoopphheerrss

	fmt.Println("---------------------")
	fmt.Println("Your turn...")

	for {
		fmt.Print("Enter a string (or type 'quit' to exit): ")
		// Reference: https://stackoverflow.com/a/34522375/8507770
		// Create a single reader which can be called multiple times
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "quit" {
			break
		}

		doubled := DoubleChars(input)
		fmt.Println(doubled)
	}
}
