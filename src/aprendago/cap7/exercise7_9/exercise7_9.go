package main

import (
	"fmt"
)

func main() {
	faveriteSport := "handball"

	switch faveriteSport {
	case "soccer":
		fmt.Printf("Today has %s\n", faveriteSport)
	case "formula1":
		fmt.Printf("Sunday will have %s\n", faveriteSport)
	case "handball":
		fmt.Printf("Saturday had %s\n", faveriteSport)
	}
}
