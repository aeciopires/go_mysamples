package main

import (
	"fmt"
)

const (
	_ = iota + 2020
	year1
	year2
	year3
	year4
)

func main() {
	fmt.Printf("%v\t%v\t%v\t%v\n", year1, year2, year3, year4)
}
