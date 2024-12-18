package main

import "fmt"

func main() {
	star := "Polaris"
	
	// Add your code below:
  // scenario 1: Declaring Pointer
  //var starAddress *string
  //starAddress = &star
  // scenario 2: Declaring Pointer
  starAddress := &star
  fmt.Println("The address of star is", starAddress)

  // dereferencing or indirecting, change value using address variable
  *starAddress = "Sirius"

  fmt.Println("The actual value of star is", star)


	greeting := "Hello there!"
	
	// Call your brainwash() below:
	brainwash(&greeting)
	
	fmt.Println("greeting is now:", greeting)
}

// Change brainwash to have a pointer parameter
func brainwash(saying *string) {
	// Dereference saying below: 
	*saying = "Beep Boop."
}
