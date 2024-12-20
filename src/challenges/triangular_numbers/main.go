/*
Reference: https://tutorialedge.net/challenges/go/triangular-numbers/
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TriangularNumber(n int) int {
	result := (n * (n + 1)) / 2
	return result
}

func main() {
	fmt.Println("*********************")
	fmt.Println("Triangular numbers are numbers that make up the sequence 1, 3, 6, 10, ...")
	fmt.Println("The nth triangular number in the sequence is the number of dots it would take to make an equilateral triangle with n dots on each side.")
	fmt.Println("The formula for the nth triangular number is (n)(n + 1) / 2")
	fmt.Println("*********************")
	fmt.Println("")

	fmt.Println("---------------------")
	fmt.Println("Your turn...")

	for {
		fmt.Print("Enter an interger number (or type 'quit' to exit): ")
		// Reference: https://stackoverflow.com/a/34522375/8507770
		// Create a single reader which can be called multiple times
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "quit" {
			break
		}

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer or 'quit'.")
			continue
		}
		fmt.Println("The sequence of", number, "triangular numbers is:")
		for i := 1; i <= number; i++ {
			result := TriangularNumber(i)
			fmt.Println(result)
		}
	}
}
