package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for aux, num := range mySlice {
		fmt.Println(aux, num)
	}

	fmt.Printf("%T\n", mySlice)
}
