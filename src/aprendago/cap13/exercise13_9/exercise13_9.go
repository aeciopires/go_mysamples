package main

import (
	"fmt"
)

// Aulas sobre callback (funcção que recebe um função como argumento):
// https://www.youtube.com/watch?v=u8qBzOAhbRM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=96
// https://www.youtube.com/watch?v=ePh12R5jnIM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=106

func sum() {
	numbers := []int{5, 4, 2, 4, 10, 15}
	total := 0
	fmt.Print("Numbers: ")
	for aux, number := range numbers {
		fmt.Print(number)
		total += number
		if aux < (len(numbers)-1){
			fmt.Print(" + ")
		}
	}
	fmt.Print(" = ")
	fmt.Print(total)
}

func sumNumbers(function func()) {
	function()
}

func main() {

	sumNumbers(sum)
}
