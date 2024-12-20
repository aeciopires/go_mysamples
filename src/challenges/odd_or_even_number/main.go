/*
Reference: https://tutorialedge.net/challenges/go/even-odd-factors/
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func OddEvenNumber(num int) string {
	var result int
	result = num % 2

	if result == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func main() {
	fmt.Println("*********************")
	fmt.Println("Odd or Even Number")
	fmt.Println("Odd numbers are numbers that cannot be divided into two equal parts.")
	fmt.Println("Even numbers are numbers that can be divided into two equal parts.")
	fmt.Println("*********************")

	var number int
	var result string

	number = 23
	result = OddEvenNumber(number)
	fmt.Println(number, "is", result) // "odd"

	number = 36
	result = OddEvenNumber(number)
	fmt.Println("36 is", result) // "even"

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
		result := OddEvenNumber(number)
		fmt.Println(number, "is", result)
	}
}
