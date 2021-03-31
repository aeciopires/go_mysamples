package main

import (
	"fmt"
)

func main() {
	// references:
	// https://blog.golang.org/strings
	// https://www.youtube.com/watch?v=AcyS0_BAy7U&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=27
	// String utf8
	s := "ascii éøâ 香"

	// Slice byte (character to character) to print: bits, position ascii, type, position and character utf8, hexa representation
	for _, v := range s {
		fmt.Printf("%b - %v - %T - %#U - %#x\n", v, v, v, v, v)
	}

	fmt.Println("")

	// Array of bytes (byte to byte) to print: bits, position ascii, type, position and character ascii, hexa representation
	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i], s[i])
	}

}
