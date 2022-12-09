package main

import (
	"fmt"
	"runtime"

	"github.com/pbnjay/memory"
)

func main() {
	fmt.Println("Number of CPU:", runtime.NumCPU())
	fmt.Printf("Total system memory: %d (bytes)\n", memory.TotalMemory())
	fmt.Printf("Free system memory: %d (bytes)\n", memory.FreeMemory())
}
