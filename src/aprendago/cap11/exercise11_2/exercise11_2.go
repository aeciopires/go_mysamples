package main

import (
	"fmt"
)

type person struct {
	name           string
	surname        string
	icecreamFlavor []string
}

func main() {
	student := make(map[string]person)

	student["Pires"] = person{
		name:           "Aecio",
		surname:        "Pires",
		icecreamFlavor: []string{"chocolate", "menta", "flocos"},
	}

	student["Pires2"] = person{
		name:           "Aecio2",
		surname:        "Pires2",
		icecreamFlavor: []string{"chocolate2", "menta2", "flocos2"},
	}

	for _, value := range student {
		fmt.Println("My name is ", value.name, " and my favorite flavors are: ")
		for _, value := range value.icecreamFlavor {
			fmt.Println("\t -", value)
		}
	}
}
