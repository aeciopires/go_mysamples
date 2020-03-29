package main

import (
	"fmt"
	"time"
)

func main() {
	// Create variables with pointers
	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	hour := &time.Time{}

	if count1 != nil {
		fmt.Printf("Address count1: %#v\n", count1)
		fmt.Printf("Value count1: %#v\n", *count1)
	}

	if count2 != nil {
		fmt.Printf("Address count2: %#v\n", count2)
		fmt.Printf("Value count2: %#v\n", *count2)
	}

	if count3 != nil {
		fmt.Printf("Address count3: %#v\n", count3)
		fmt.Printf("Value count3: %#v\n", *count3)
	}

	if hour != nil {
		fmt.Printf("Address hour: %#v\n", hour)
		fmt.Printf("Value hour: %#v\n", *hour)
		fmt.Printf("Value hour: %#v\n", hour.String())
	}
}
