package main

// references:
// https://www.youtube.com/watch?v=USntMXkOihY&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=31
// https://splice.com/blog/iota-elegant-constants-golang/
// https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827

import (
	"fmt"
)

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	fmt.Println("binary\t\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}
