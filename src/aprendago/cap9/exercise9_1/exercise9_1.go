package main

import (
	"fmt"
)

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}

	for aux, num := range myArray {
		fmt.Println(aux, num)
	}

	fmt.Printf("%T\n", myArray)
}
