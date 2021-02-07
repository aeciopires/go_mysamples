package main

import (
	"fmt"
)

var endYear int = 2021

func main() {
	year := 1986

	for {
		if year > endYear {
			break
		}
		fmt.Printf("Year: %d\n", year)
		year++
	}
}
