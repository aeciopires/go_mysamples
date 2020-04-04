package main

import "fmt"

// OBS 1.: When you define an array, you must specify what type of data it may contain and how big the array is in the following form: [<size>]<type>. The key to making this an array is specifying the size. If your definition didn't have the size, it would seem like it works, but it would not be an array – it'd be a slice.

// OBS 2.: The array's length is part of its type definition. If you have two arrays that accept the same type but they're different sizes, they are not compatible and aren't comparable with each other. Arrays of different lengths that are not the same type can can't be compared with each other.

// Reference: https://courses.packtpub.com/courses/take/go/texts/9761782-exercise-4-01-defining-an-array

func defineArray() [10]int {
	// Array will be initialized with element default of int
	var arr [10]int
	return arr
}

func main() {
	fmt.Printf("%#v\n", defineArray())
}
