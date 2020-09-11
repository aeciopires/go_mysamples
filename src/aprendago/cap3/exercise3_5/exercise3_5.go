package main

import (
	"fmt"
)

type mytype int

var x mytype
var y int

func main() {
	fmt.Printf("Before x=%v (%T)\n", x, x)
	x = 42
	fmt.Printf("After x=%v (%T)\n", x, x)
	y = int(x)
	fmt.Printf("y=%v (%T)", y, y)
}
