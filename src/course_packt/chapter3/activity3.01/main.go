package main

import (
	"fmt"
)

func calculateTax(costItem []float32, taxRate []float32) float32 {
	var taxTotal float32 = 0
	for aux := 0; aux < len(costItem); aux++ {
		taxTotal = taxTotal + ((costItem[aux] * taxRate[aux]) / 100)
		//fmt.Println("Sub taxTotal", taxTotal)
	}
	return taxTotal
}

func main() {
	cost := []float32{0.99, 2.75, 0.87}
	tax := []float32{7.5, 1.5, 2}
	var total float32 = 0

	total = calculateTax(cost, tax)
	fmt.Println("Sales Tax Total", total)
}
