package main

import (
	"fmt"
)

func main() {
	//               0   1   2   3   4   5   6   7   8   9
	mySlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	mySlice2 := append(mySlice[:3], mySlice[len(mySlice)-4:]...)
	fmt.Println(mySlice2)
}
