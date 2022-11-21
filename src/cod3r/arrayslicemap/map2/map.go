package main

import "fmt"

func main() {
	applications := map[string]string{
		"gateway": "1.0.1",
		"manager": "2.5.0",
		"auth":    "0.10.0",
	}

	fmt.Println(applications)
	applications["front"] = "0.1.0"
	fmt.Println(applications)
	delete(applications, "database") // this element not exists in map, but an error ins't happen
	fmt.Println(applications)

	for name, version := range applications {
		fmt.Println(name, version)
	}
}
