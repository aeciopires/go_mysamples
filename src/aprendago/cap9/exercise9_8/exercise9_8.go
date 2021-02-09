package main

import (
	"fmt"
)

func main() {

	person := map[string][]string{
		"junior_joao": []string{
			"andar", "correr",
		},
		"pereira_ze": []string{
			"pular", "voar",
		},
	}

	for key, value := range person {
		fmt.Println(key)
		for i, hobby := range value {
			fmt.Println("\t", i, " - ", hobby)
		}
	}

	fmt.Println(person)
}
