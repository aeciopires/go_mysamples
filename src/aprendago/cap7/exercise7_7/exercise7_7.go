package main

import (
	"fmt"
)

func main() {
	fruit1 := "banana"
	fruit2 := "apple"
	number := 4
	endNumber := 100

	for {
		if number <= endNumber {
			module := number % 4
			if module >= 3 {
				fmt.Printf("Number: %d Module by 4: %d\tFruit: %s\n", number, module, fruit1)
			} else {
				fmt.Printf("Number: %d Module by 4: %d\tFruit: %s\n", number, module, fruit2)
			}
		} else {
			break
		}
		number++
	}
}
