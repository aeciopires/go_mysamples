package main

import "fmt"

func message() string {
	arr := [...]string{
		"ready",
		"Get",
		"Go",
		"to",
	}
	// OBS.: Create a message by joining the words in a specific order and returning it. We're using the fmt.Sprintln function here since it allows us to capture the formatted text before it's printed.

	// Reference: https://courses.packtpub.com/courses/take/go/texts/9761802-exercise-4-04-reading-a-single-item-from-an-array
	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}

func main() {
	fmt.Print(message())
}
