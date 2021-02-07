package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"apple", "banana", "grape"}

	mySlice2 := mySlice[:]
	fmt.Println(mySlice2)

	fmt.Println()

	for i := 0; i < len(mySlice); i++ {
		fmt.Println(mySlice2[i])
	}

	fmt.Println()

	mySlice3 := append(mySlice[:1], mySlice[2:]...)
	fmt.Println(mySlice3)
}
