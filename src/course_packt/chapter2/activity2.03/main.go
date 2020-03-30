// Reference: https://tutorialedge.net/golang/implementing-bubble-sort-with-golang/
package main

import (
	"fmt"
)

func myBubbleSort(listNumbers []int) {
	// set swapped to true
	swapped := true

	for swapped {
		// set swapped to false
		swapped = false
		// iterate through all of the elements in list
		for aux := 1; aux < len(listNumbers); aux++ {
			// if the current element is greater than the next
			// element, swap them
			if listNumbers[aux-1] > listNumbers[aux] {
				fmt.Println("Swapping element")
				// swap values using Go's tuple assignment
				listNumbers[aux], listNumbers[aux-1] = listNumbers[aux-1], listNumbers[aux]
				// set swapped to true - this is important
				// if the loop ends and swapped is still equal
				// to false, the algorithm will assume the list is
				// fully sorted.
				swapped = true
			}
		}
	}
}

func main() {
	numsBefore := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	fmt.Println("Before:", numsBefore)
	myBubbleSort(numsBefore)
	fmt.Println("After:", numsBefore)
}
