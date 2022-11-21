package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber() int {
	maxInterval := 1000

	sourceNumbers := rand.NewSource(time.Now().UnixNano()) // generate a group of numbers in nanoseconds format
	//fmt.Println(sourceNumbers)
	//fmt.Println("------------------")

	randomInterger := rand.New(sourceNumbers) // select a random number of the group of numbers in nanoseconds format
	//fmt.Println(randomInterger)
	//fmt.Println("------------------")

	return randomInterger.Intn(maxInterval)
}

func main() {
	threshold := 100
	if i := randomNumber(); i > threshold { // this is too support in switch statement
		fmt.Println("The random number:", i, "is major of than:", threshold)
	} else {
		fmt.Println("The random number:", i, "is minor of than:", threshold)
	}
}
