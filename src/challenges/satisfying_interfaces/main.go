/*
Reference: https://tutorialedge.net/challenges/go/type-assertions/
*/

package main

// Employee interface
type Employee interface {
  // Interface methods
  Language() string
  Age() int
}

// Setting a global struct
type Engineer struct {
  Name string
}

// Function that receive a struct Engineer and
// implements Language() method of interface
func (e Engineer) Language() string {
  return e.Name + " programs in Go"
}

// Function that receive a struct Engineer and
// implements Age() method of interface
func (e Engineer) Age() int {
  return 26
}

func main() {
  // This will throw an error
  var programmers []Employee
  elliot := Engineer{Name: "Elliot"}
  // Engineer does not implement the Employee interface
  // you'll need to implement Age() and Random()
  programmers = append(programmers, elliot)
}
