package main

import (
	"fmt"
)

// Typed const
const x int = 20

// Untyped const
const y = 10

func main() {
	fmt.Printf("Typed const: %v (%T)\n", x, x)
	fmt.Printf("Untyped const: %v (%T)\n", y, y)
}
