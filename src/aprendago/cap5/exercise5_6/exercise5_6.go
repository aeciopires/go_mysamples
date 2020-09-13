package main

import (
	"fmt"
)

const (
	_    = iota + 2020
	ano1
	ano2
	ano3
	ano4
)

func main() {
	fmt.Printf("%v\t%v\t%v\t%v\n", ano1, ano2, ano3, ano4)
}