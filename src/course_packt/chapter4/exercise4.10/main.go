package main

import (
	"fmt"
)

func message() string {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	/*
		We'll extract the first value, first directly as an int, then as a slice using both low and high, and finally using just high and skipping low. We'll write the values to a message string
	*/
	m := fmt.Sprintln("First :", s[0], s[0:1], s[:1])

	/*
		We'll get the last element. To get the int, we'll use the length and subtract 1 from the index. We use that same logic when setting the low for the range notation. For high, we can use the length of the slice. Finally, we can see we can skip high and get the same result
	*/
	m += fmt.Sprintln("Last  :", s[len(s)-1], s[len(s)-1:len(s)], s[len(s)-1:])

	// Let's get the first five values and add them to the message
	m += fmt.Sprintln("First 5 :", s[:5])

	// We'll get the last four values and add them to the message as well
	m += fmt.Sprintln("Last 4 :", s[5:])

	// Finally, we'll extract five values from the middle of the slice and get them in the message too
	m += fmt.Sprintln("Middle 5:", s[2:7])

	return m
}

func main() {
	fmt.Print(message())
}
