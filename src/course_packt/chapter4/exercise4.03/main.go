package main

import "fmt"

func compArrays() (bool, bool, [10]int) {
	var arr1 [10]int
	arr2 := [...]int{9: 0}
	arr3 := [10]int{1, 9: 10, 4: 5}
	return arr1 == arr2, arr1 == arr3, arr3
}

// OBS.: We used keys when initializing the data for an array. For arr2, we combined the ... shortcut with setting a key to make the array length directly relate to the key we set. With arr3, we mixed it using keys and without using keys, and we also used the keys out of order. Go's flexibility when using keys is strong and makes using arrays in this way pleasant.

// Reference: https://courses.packtpub.com/courses/take/go/texts/9761795-exercise-4-03-initializing-an-array-using-keys

func main() {
	comp1, comp2, arr3 := compArrays()
	fmt.Println("[10]int == [...]{9:0}       :", comp1)
	fmt.Println("[10]int == [10]int{1, 9: 10, 4: 5}}:", comp2)
	fmt.Println("arr3               :", arr3)
}
