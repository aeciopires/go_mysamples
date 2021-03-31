package main

import (
	"fmt"
)

type person struct {
	name    string
	surname string
	age     int
}

func (p person) showPerson() {
	fmt.Println("Complete name:", p.name, p.surname)
	fmt.Println("Age:", p.age)
}

func main() {
	son := person{
		name:    "Aecio",
		surname: "Pires",
		age:     17,
	}
	son.showPerson()
}
