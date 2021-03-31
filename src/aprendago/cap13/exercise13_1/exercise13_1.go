package main

import (
	"fmt"
	"math/rand"
)

//func (receiver) identifier(parameters) (returns) { code }

func returnNumberInteger() int {
	return rand.Intn(100)
}

func returnNumberIntegerString() (int, string) {
	number := returnNumberInteger()
	word := "Congratulations by X years old."
	return number, word
}

func main() {
	fmt.Println(returnNumberInteger())
	fmt.Println(returnNumberIntegerString())
}
