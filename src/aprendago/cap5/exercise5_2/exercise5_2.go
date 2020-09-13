package main

import (
	"fmt"
)

func main() {
	a := 10 == 10
	b := 10 != 20
	c := 20 <= 20
	d := 10 < 20
	e := 21 >= 20
	f := 20 > 10

	fmt.Printf("(10 == 10) Equal? %v\n", a)
	fmt.Printf("(10 != 20) Diferent? %v\n", b)
	fmt.Printf("(20 <= 20) Less Than Equal To? %v\n", c)
	fmt.Printf("(10 < 20)  Less Than? %v\n", d)
	fmt.Printf("(21 >= 20) Greater Than Equal To? %v\n", e)
	fmt.Printf("(20 > 10)  Greater Than? %v\n", f)
}
