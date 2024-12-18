package main

import "fmt"

// Define logData Here
func logData(data any){
  fmt.Println("Data logged: ", data)
}

func main() {
    // Call logData with a string here
    logData("pass")
    // Call logData with an integer here
    logData(2)
}
