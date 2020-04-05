package main

import "fmt"

func getElements() []string {
	elements := []string{"good", "good", "bad", "good", "good"}
	// Create a range that starts at index 0 and goes to index 2.
	// Then, create a slice range that starts at the index 3 of the slice and goes up to index 4.
	// Use append to add the second range to the first range. Capture the value from append
	elements = append(elements[0:2], elements[3:]...)
	return elements
}

func main() {
	fmt.Println(getElements())
}
