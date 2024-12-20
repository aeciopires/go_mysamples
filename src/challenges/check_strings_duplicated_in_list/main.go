/*
Reference: https://tutorialedge.net/challenges/go/checking-for-duplicates/
*/

package main

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	uniqueNames := make(map[string]bool)
	var unique []string
	for _, dev := range developers {
		if _, ok := uniqueNames[dev.Name]; !ok {
			unique = append(unique, dev.Name)
			uniqueNames[dev.Name] = true
		}
	}

	return unique 
}

func main() {
	var employees = []Developer{
		Developer{Name: "Elliot", Age: 28},
		Developer{Name: "Alan", Age: 29},
		Developer{Name: "Jennifer", Age: 21},
		Developer{Name: "Graham", Age: 18},
		Developer{Name: "Paul", Age: 40},
		Developer{Name: "Alan", Age: 29},
	}

	fmt.Println("Filter Unique Challenge")
	fmt.Println()
	fmt.Println("Original list:")
	fmt.Println(employees)
	fmt.Println("---------------------")
	result := FilterUnique(employees)
	fmt.Println("Filtered list:")
	fmt.Println(result)
}
