package main

import (
	"fmt"
)

func main() {
	fruit0 := "grape"
	fruit1 := "orange"
	fruit2 := "apple"
	fruit3 := "banana"
	number := 4
	endNumber := 100
	constModule := 4
	number2 := 5
	constModule2 := 5
	endNumber2 := 50

	// Example of statement switch undefined
	for {
		if number <= endNumber {
			module := number % constModule
			switch {
			case module == 1:
				fmt.Printf("Number: %d Module by %d: %d\tFruit: %s\n", number, constModule, module, fruit1)
			case module == 2:
				fmt.Printf("Number: %d Module by %d: %d\tFruit: %s\n", number, constModule, module, fruit2)
			case module == 3:
				fmt.Printf("Number: %d Module by %d: %d\tFruit: %s\n", number, constModule, module, fruit3)
			default:
				break
			}
			number++
		} else {
			break
		}
	}

	fmt.Println()

	// Example of statement switch defined
	for {
		if number2 <= endNumber2 {
			module2 := number2 % constModule2
			switch module2 {
			case 0:
				fmt.Printf("Number: %d Module by %d: %d\tFruit: %s\n", number2, constModule2, module2, fruit0)
			case 1:
				fmt.Printf("Number: %d Module by %d: %d\tFruit: %s\n", number2, constModule2, module2, fruit1)
			case 2:
				fmt.Printf("Number: %d Module by %d: %d\tFruit: %s\n", number2, constModule2, module2, fruit2)
			case 3:
				fmt.Printf("Number: %d Module by %d: %d\tFruit: %s\n", number2, constModule2, module2, fruit3)
			default:
				break
			}
			number2++
		} else {
			break
		}
	}
}
