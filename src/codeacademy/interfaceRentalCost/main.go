package main

import "fmt"

// PrintRentalCost prints the cost of renting a vehicle 
func PrintRentalCost(v Vehicle) { 
  fmt.Printf("Rental Cost: $%d\n", v.CalculateRentalCost())
}

func main() {
  var vehicle Vehicle
  //vehicle = Scooter{hours: 4}
  vehicle = Car{hours: 2, miles: 4}

	PrintRentalCost(vehicle)
}

type Vehicle interface { // Interface defining methods required for vehicles
  CalculateRentalCost() int
}

type Scooter struct { // Scooter struct for storing scooter usage data
  hours int // Hours the scooter has been used
}

func (s Scooter) CalculateRentalCost() int {
  const hourlyRate int = 5 // Scooter hourly rate
  return s.hours * hourlyRate
}

type Car struct { // Car struct for storing car usage data
  hours int // Hours the car has been used
  miles int // Miles the car has been driven
}

func (c Car) CalculateRentalCost() int {
  const hourlyRate int = 10 // Car hourly rate
  const mileageRate int = 5 // Car mileage rate
  return (c.hours * hourlyRate) + (c.miles * mileageRate)
}
