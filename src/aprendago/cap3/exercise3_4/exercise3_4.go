package main

// Reference: https://golang.org/ref/spec#Types

import (
	"fmt"
)

type mytype int

var x mytype

func main() {
	fmt.Printf("Before x=%v (%T)\n", x, x)
	x = 42
	fmt.Printf("After x=%v (%T)", x, x)
}
