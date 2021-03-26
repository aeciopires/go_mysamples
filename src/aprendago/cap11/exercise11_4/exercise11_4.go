package main

import (
	"fmt"
)

func main() {
	car := struct{
		name string
		year []int
		details map[string]string
	}{
		name: "challenge",
		year: []int{1970, 1971,},
		details: map[string]string{
			"color": "blue",
		},
	}

	fmt.Println(car)
}
