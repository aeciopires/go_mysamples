package main

import (
	"fmt"
)

// Aulas sobre clousure (função que retorne um função que faz uso de uma variável além do seu escopo):
// https://www.youtube.com/watch?v=mOM0qTB5ppU&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=97
// https://www.youtube.com/watch?v=LDJF2ceCRDE&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=107
// https://github.com/ellenkorbes/aprendago/issues/50

func sum() func() int {
	number1 := 1
	number2 := 2
	return func() int {
		number1++
		number2++
		return number1 + number2
	}
}

func main() {
	total := sum()
	fmt.Println(total())
	fmt.Println(total())
	fmt.Println(total())
	fmt.Println(total())
	fmt.Println(total())
}