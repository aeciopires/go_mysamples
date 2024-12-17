package main

import "fmt"

func main() {
    // Scenario 1
    //discount := NoDiscount{}
    // Scenario 2
    //discount := PercentageDiscount{percent: 10}
    // Scenario 3
    //discount := FixedAmountDiscount{amount: 20}
    store := Store{discountType: discount}
	
    initialPrice := 100.0
	  store.printDiscountedPrice(initialPrice)
}

// Define the interface
type Discount interface {
    // Define method signature of the interface
    Apply(amount float64) float64
}

// Implementation 1: PercentageDiscount
type PercentageDiscount struct {
    percent float64
}

func (d PercentageDiscount) Apply(amount float64) float64 {
    return amount - (amount * d.percent / 100)
}

// Implementation 2: FixedAmountDiscount
type FixedAmountDiscount struct {
    amount float64
}

func (d FixedAmountDiscount) Apply(amount float64) float64 {
    return amount - d.amount
}

// Implementation 3: NoDiscount
type NoDiscount struct{}

func (d NoDiscount) Apply(amount float64) float64 {
    return amount
}

// Store struct that delegates to Discount 
type Store struct{
	discountType Discount
} 

// Function to use the Discount interface
func (s Store) printDiscountedPrice(amount float64) {
    fmt.Printf("Discounted Price: %.2f\n", s.discountType.Apply(amount))
}
