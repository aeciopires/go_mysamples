package main

import "fmt"

func main() {
  // Receiver of interface
	var vehicle Vehicle
  // Setting type Car
	vehicle = Car{ hours: 2, miles: 4 }
  // calling function
	PrintRentalCost(vehicle)
  // Setting type Scooter
	vehicle = Scooter{ hours: 3 }
  // calling function
	PrintRentalCost(vehicle)
  // Setting type Bike
	vehicle = Byke{ hours: 1 }
  // calling function
	PrintRentalCost(vehicle)
}

// PrintRentalCost prints the cost of renting a vehicle, calling the interface method: v.CalculateRentalCost
func PrintRentalCost(v Vehicle) { 
	fmt.Printf("Rental Cost: $%d\n", v.CalculateRentalCost())
}

type Vehicle interface { // Interface defining methods required for vehicles
	CalculateRentalCost() int
}

type Scooter struct { // Scooter struct for storing scooter usage data
	hours int // Hours the scooter has been used
}

// Function to implement the interface method Vehicle.CalculateRentalCost()
func (s Scooter) CalculateRentalCost() int {
	const hourlyRate int = 5 // Scooter hourly rate
	return s.hours * hourlyRate
}

type Car struct { // Car struct for storing car usage data
	hours int // Hours the car has been used
	miles int // Miles the car has been driven
}

// Function to implement the interface method Vehicle.CalculateRentalCost()
func (c Car) CalculateRentalCost() int {
	const hourlyRate int = 10 // Car hourly rate
	const mileageRate int = 5 // Car mileage rate
	return (c.hours * hourlyRate) + (c.miles * mileageRate)
}

type Byke struct { // Bike struct for storing bike usage data
	hours int // Hours the car has been used
}

// Function to implement the interface method Vehicle.CalculateRentalCost()
func (b Byke) CalculateRentalCost() int {
	const hourlyRate int = 2 // Byke hourly rate
	return (b.hours * hourlyRate)
}
