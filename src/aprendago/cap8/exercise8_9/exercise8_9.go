package main

import (
	"fmt"
)

func main() {

	routers := map[string]string{
		"cisco":    "c3640",
		"mikrotik": "rb1100",
	}

	fmt.Println(routers)
	fmt.Println(routers["cisco"])

	routers["junniper"] = "srx300"

	fmt.Println(routers)
	fmt.Println(routers["junniper"], "\n\n")

	// comma ok idiom
	if exist, ok := routers["arista"]; !ok {
		fmt.Println("[ERROR] Not found!")
	} else {
		fmt.Println(exist)
	}
}
