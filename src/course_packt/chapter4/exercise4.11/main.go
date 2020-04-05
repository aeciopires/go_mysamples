package main

import "fmt"

func genSlices() ([]int, []int, []int) {
	var s1 []int

	// Define a slice using make and set only the length
	s2 := make([]int, 10)

	// Define a slice that uses both the length and capacity of the slices
	s3 := make([]int, 10, 50)

	return s1, s2, s3
}

func main() {
	/*
		Due to the complexity of what a slice is and how it works, you can't compare slices to one another. If you try, Go gives you an error. You can compare a slice to nil, but that's it.

		A slice is not a value, and it's not a pointer, so what is it? A slice is a special construct in Go. A slice doesn't store its own values directly. In the background, it's using an array that you can't access directly. What a slice does store is a pointer to that hidden array, its own starting point in that array, how long the slice is, and what the capacity of the slice is. These values provide slices with a window for the hidden array.

		Reference: https://courses.packtpub.com/courses/take/go/texts/9762081-exercise-4-12-controlling-internal-slice-behavior
	*/

	s1, s2, s3 := genSlices()
	fmt.Printf("s1: len = %v cap = %v\n", len(s1), cap(s1))
	fmt.Printf("s2: len = %v cap = %v\n", len(s2), cap(s2))
	fmt.Printf("s3: len = %v cap = %v\n", len(s3), cap(s3))
}
