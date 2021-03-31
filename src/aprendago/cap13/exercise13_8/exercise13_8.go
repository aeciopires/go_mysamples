package main

import (
	"fmt"
)

// Aulas sobre funções que retornam função:
// https://www.youtube.com/watch?v=9Oxmya_A-Sc&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=95
// https://www.youtube.com/watch?v=ePh12R5jnIM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=105

func sumNumbers() func() {
	return func() {
		numbers := []int{5, 4, 2, 4, 10, 15}
		total := 0
		fmt.Print("Numbers: ")
		for aux, number := range numbers {
			fmt.Print(number)
			total += number
			if aux < (len(numbers) - 1) {
				fmt.Print(" + ")
			}
		}
		fmt.Print(" = ")
		fmt.Print(total)
	}
}

func main() {

	sum := sumNumbers()
	sum()
}
