package main

import (
	"fmt"
)

func main() {
	endNumber := 100
	number := 10

	for {
		if number <= endNumber {
			module := number % 4
			fmt.Printf("Number: %d Module by 4: %d\n", number, module)
		} else {
			break
		}
		number++
	}
}
