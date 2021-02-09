package main

import (
	"fmt"
)

func main() {
	//               0  1  2  3  4  5  6  7  8  9
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mySlice2 := append(mySlice[0:3])
	fmt.Println(mySlice2)

	mySlice3 := append(mySlice[4:])
	fmt.Println(mySlice3)

	mySlice4 := append(mySlice[1:7])
	fmt.Println(mySlice4)

	mySlice5 := append(mySlice[2 : len(mySlice)-1])
	fmt.Println(mySlice5)
}
