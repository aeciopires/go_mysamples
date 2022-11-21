package main

import "fmt"

func main() {
	// var applications map[string]string

	// maps must be initialized
	applications := make(map[string]string)

	applications["gateway"] = "1.0.1"
	applications["manager"] = "2.5.0"
	applications["auth"] = "0.10.1"
	fmt.Println(applications)

	for name, version := range applications {
		fmt.Printf("%s (Version: %s)\n", name, version)
	}

	fmt.Println("Version of gateway:", applications["gateway"])
	delete(applications, "gateway")

	fmt.Println(applications)
}
