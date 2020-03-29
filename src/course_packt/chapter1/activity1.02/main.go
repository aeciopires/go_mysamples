package main

import "fmt"

func main() {
	myvar1, myvar2 := 5, 10

	// Call swap function passing address pointers of variables
	swap(&myvar1, &myvar2)
	fmt.Println(myvar1 == 10, myvar2 == 5)
}

func swap(var1 *int, var2 *int) {
	// Change values in pointers of variables
	*var1 += 5
	*var2 -= 5
}
