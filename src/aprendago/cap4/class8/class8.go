package main

// reference: https://www.youtube.com/watch?v=1IduyaGMO3g&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=30

import (
	"fmt"
)

const (
	// Numa declaração de constantes, o identificador iota representa números sequenciais
	a = iota + 10000000
	_ //descarta o valor 10000001
	c
	x
	_ //descarta o valor 10000004
	z
)

func main() {
	fmt.Println(a, c, x, z)
}
