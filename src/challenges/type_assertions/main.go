/*
Reference: https://tutorialedge.net/challenges/go/type-assertions/
*/

package main

import (
  "fmt"
)

// Setting a global struct
type Developer struct {
  Name string
  Age int
}

// Function that receive two dynamic interface as argument and return a struct
func GetDeveloper(name interface{}, age interface{}) Developer {
  // Setting a local variable of struct type
  var dev Developer

  // Setting properties using the interfaces
  dev.Name = name.(string)
  dev.Age = age.(int)

  // Returning the struct
  return dev
}

func main() {
  fmt.Println("Hello World")

  // Setting two variables using dynamic interfaces
  var name interface{} = "Elliot"
  var age interface{} = 26

  // Calling the function
  dynamicDev := GetDeveloper(name, age)
  fmt.Println("My name is:", dynamicDev.Name, ". I am", dynamicDev.Age, "years old")
}
