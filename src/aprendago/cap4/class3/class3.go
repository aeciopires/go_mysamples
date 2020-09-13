package main

import (
	"fmt"
	"runtime"
)

func main() {
	// show operating system. Ex.: linux
	fmt.Println(runtime.GOOS)
	// show processor architecture. Ex.: amd64
	fmt.Println(runtime.GOARCH)
}
