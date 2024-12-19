/*
Reference: https://tutorialedge.net/challenges/go/sort-by-price/
*/

package main

import (
	"fmt"
	"sort"
)

// Flight - a struct that
// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

// Setting an slice of Flight
type ByPrice []Flight

// This code snippet defines a method named Len() on the ByPrice type.
// ByPrice is a custom type defined as a slice of Flight structs: type ByPrice []Flight.
// The Len() method returns the number of elements in the underlying slice p.
// This method is one of the three methods required to satisfy the sort.Interface interface in Go's sort package.
// By implementing Len(), Swap(), and Less(), you enable the use of the sort.Sort() function on your custom data types.
func (p ByPrice) Len() int {
	return len(p)
}

// The Swap(i, j int) method swaps the elements at indices i and j within the underlying slice p of Flight structs.
// This operation is essential for the sorting algorithm to rearrange elements into the desired order.
// The concise syntax p[i], p[j] = p[j], p[i] utilizes Go's simultaneous assignment feature to efficiently perform the swap without needing a temporary variable.
func (p ByPrice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// The Less(i, j int) method compares the Price field of two Flight structs at indices i and j within the ByPrice slice p.
// It returns true if the price of the element at index i is greater than the price of the element at index j (p[i].Price > p[j].Price).
func (p ByPrice) Less(i, j int) bool {
	return p[i].Price > p[j].Price
}

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights []Flight) []Flight {
	sort.Sort(ByPrice(flights))
	return flights
}

// Show results
func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d | ", flight.Origin, flight.Destination, flight.Price)
	}
}

func main() {
	// Setting slice
	flights := []Flight{
		Flight{Origin: "VCP", Destination: "DST1", Price: 30},
		Flight{Origin: "LAX", Destination: "DST2", Price: 20},
		Flight{Origin: "JPA", Destination: "DST3", Price: 50},
		Flight{Origin: "GRU", Destination: "DST4", Price: 1000},
	}

	sort.Sort(ByPrice(flights))

	printFlights(flights)
}
