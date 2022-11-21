package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Reference: https://stackoverflow.com/a/34522375/8507770
	// Create a single reader which can be called multiple times
	reader := bufio.NewReader(os.Stdin)
	// Prompt and read
	fmt.Print("Enter first number (type integer): ")
	string1, _ := reader.ReadString('\n')
	fmt.Print("Enter second number (type integer): ")
	string2, _ := reader.ReadString('\n')

	number1, _ := strconv.Atoi(strings.TrimSpace(string1))
	number2, _ := strconv.Atoi(strings.TrimSpace(string2))
	fmt.Printf("[DEBUG] Number1: \"%d\", Number2: \"%d\"\n", number1, number2)

	result := sum2IntegerNumbers(number1, number2)
	printResultIntegerSum(result)
}
