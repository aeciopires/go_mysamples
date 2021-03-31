package main

import (
	"fmt"
)

// Aulas sobre funções atribuídas a uma variável:
// https://www.youtube.com/watch?v=j9C66R4BMWM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=94
// https://www.youtube.com/watch?v=ePh12R5jnIM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=104

func main() {

	sumNumbers := func() {
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

	sumNumbers()
}
