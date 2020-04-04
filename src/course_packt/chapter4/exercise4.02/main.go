package main

import "fmt"

// OBS.: The array's length is part of its type definition. If you have two arrays that accept the same type but they're different sizes, they are not compatible and aren't comparable with each other. Arrays of different lengths that are not the same type can can't be compared with each other.

// You should see an error. This error is telling you that arr1, which is a [5] int, and arr4, which is a [9] int, are not the same type and aren't compatible. Let's fix that.

// Reference: https://courses.packtpub.com/courses/take/go/texts/9761782-exercise-4-01-defining-an-array

func compArrays() (bool, bool, bool) {
	var arr1 [5]int
	arr2 := [5]int{0}
	arr3 := [...]int{0, 0, 0, 0, 0}
	arr4 := [9]int{0, 0, 0, 0, 9}
	return arr1 == arr2, arr1 == arr3, arr1 == arr4
}

func main() {
	comp1, comp2, comp3 := compArrays()
	fmt.Println("[5]int == [5]int{0}       :", comp1)
	fmt.Println("[5]int == [...]int{0, 0, 0, 0, 0}:", comp2)
	fmt.Println("[5]int == [9]int{0, 0, 0, 0, 9} :", comp3)
}
