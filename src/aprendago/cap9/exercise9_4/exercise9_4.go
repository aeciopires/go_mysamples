package main

import (
	"fmt"
)

func main() {
	//               0   1   2   3   4   5   6   7   8   9
	mySlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	mySlice = append(mySlice, 52)
	fmt.Println(mySlice)

	mySlice = append(mySlice, 53, 54, 55)
	fmt.Println(mySlice)

	//                0   1   2   3   4
	mySlice2 := []int{56, 57, 58, 59, 60}
	mySlice = append(mySlice, mySlice2...)
	fmt.Println(mySlice)
}
