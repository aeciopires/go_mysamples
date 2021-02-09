package main

import (
	"fmt"
)

func main() {
	// multi-dimensional slice of string
	table := [][]string{
		//Colunn:     0       1       2          // Row:
		[]string{"joao", "pereira", "andar"},    //  0
		[]string{"joaquim", "nabuco", "correr"}, //  1
		[]string{"jonas", "prestes", "pular"},   //  2
	}

	for _, i := range table {
		fmt.Println(i)
		for _, j := range i {
			fmt.Printf("%s\t\n", j)
		}
		fmt.Println("")
	}
}
