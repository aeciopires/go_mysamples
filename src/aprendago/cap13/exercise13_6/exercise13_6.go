package main

import (
	"fmt"
)

// Aulas sobre funções anônimas:
// https://www.youtube.com/watch?v=pp8NKaoyefQ&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=93
// https://www.youtube.com/watch?v=NAXWCqhJIEU&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=103


func main() {

	func() {
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
	}()
}
