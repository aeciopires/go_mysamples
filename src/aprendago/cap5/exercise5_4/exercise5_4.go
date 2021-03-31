package main

import (
	"fmt"
)

func main() {
	x := 259
	fmt.Printf("Decimal: %d\nBinário: %b\nHexadecimal: %#x\n", x, x, x)
	// Bits shifted to the left
	y := x << 1
	fmt.Printf("Decimal: %d\nBinário: %b\nHexadecimal: %#x\n", y, y, y)
}
