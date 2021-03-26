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
	student1 := person{
		name:           "Aecio",
		surname:        "Pires",
		icecreamFlavor: []string{"chocolate", "menta", "flocos"},
	}

	student2 := person{
		name:           "Aecio2",
		surname:        "Pires2",
		icecreamFlavor: []string{"chocolate2", "menta2", "flocos2"},
	}

	fmt.Println(student1.name, student1.surname)
	for _, value := range student1.icecreamFlavor {
		fmt.Println("\t", value)
	}

	fmt.Println(student2.name, student2.surname)
	for _, value := range student2.icecreamFlavor {
		fmt.Println("\t", value)
	}
}
