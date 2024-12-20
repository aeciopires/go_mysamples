/*
Reference: https://tutorialedge.net/challenges/go/calculate-armstrong-number/
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyInt int

func (n *MyInt) IsArmstrong() bool {
	num := int(*n)
	sum := 0
	temp := num

	fmt.Println("Show details...")

	for temp > 0 {
		digit := temp % 10
		sum += digit * digit * digit
		fmt.Println("Digit:", digit)
		fmt.Println("Digit to cube:", digit*digit*digit)
		temp /= 10
	}

	fmt.Println("Result of sum of cube: ", sum)
	fmt.Println("")
	return num == sum
}

func showResults(number MyInt) {
	fmt.Println("")
	fmt.Println("The number is: ", number)
	fmt.Println("")
	fmt.Println("Is it an Armstrong number? ", number.IsArmstrong())
	fmt.Println("")
}

func main() {
	fmt.Println("*********************")
	fmt.Println("Armstrong Numbers")
	fmt.Println("An Armstrong number is a 3-digit number such that each of the sum of the cubes of its digits equal the number itself")
	fmt.Println("*********************")

	var num1 MyInt = 371
	showResults(num1)

	var num2 MyInt = 369
	showResults(num2)

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

		num3Int, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer or 'quit'.")
			continue
		}
		num3 := MyInt(num3Int)
		showResults(num3)
	}
}
