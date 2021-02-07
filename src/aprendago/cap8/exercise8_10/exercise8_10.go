package main

import (
	"fmt"
)

func main() {

	routers := map[int]string{
		1: "c3640",
		2: "rb1100",
	}

	fmt.Println(routers)

	total := 0

	for key := range routers {
		total += key
	}

	fmt.Println(total)

	fmt.Println(routers)
	delete(routers, 2)
	fmt.Println(routers)
}
