package main

import (
	"fmt"
)

func main() {
	for letter := 65; letter <= 90; letter++ {
		fmt.Printf("Number: %d\n", letter)
		for aux := 1; aux <= 3; aux++ {
			fmt.Printf("\tUnicode: %U\tLetter: %c\n", letter, letter)
		}
	}
}
