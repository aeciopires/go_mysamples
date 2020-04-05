package main

import "fmt"

func getWeek() []string {
	week := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	// Create a range that starts at index 6 and goes to the end of the slice.
	// Then, create a slice range that starts at the beginning of the slice and goes up to index 6.
	// Use append to add the second range to the first range. Capture the value from append
	week = append(week[6:], week[:6]...)
	return week
}

func main() {
	fmt.Println(getWeek())
}
