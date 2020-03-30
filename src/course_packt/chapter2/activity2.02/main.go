package main

import "fmt"

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	auxCount := 0
	auxWord := ""

	for word, count := range words {
		if count > auxCount {
			auxCount = count
			auxWord = word
		}
	}

	fmt.Println("Most popular word: ", auxWord)
	fmt.Println("With a count of: ", auxCount)
}
