/*
Reference: https://tutorialedge.net/challenges/go/leap-years/ (anos bissextos)
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CheckLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		return false
	}
}

func main() {
	fmt.Println("Check Leap Year")
	fmt.Println("*********************")
	fmt.Println("The formula to determine if a year is a leap year is:")
	fmt.Println("Divisible by 4: Years that are evenly divisible by 4 are leap years.")
	fmt.Println("Not divisible by 100: Years that are divisible by 100 are not leap years.")
	fmt.Println("Divisible by 400: Years that are divisible by 100 and 400 are leap years.")
	fmt.Println("*********************")

	fmt.Println("---------------------")
	fmt.Println("Your turn...")

	for {
		fmt.Print("Enter a year (or type 'quit' to exit): ")
		// Reference: https://stackoverflow.com/a/34522375/8507770
		// Create a single reader which can be called multiple times
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "quit" {
			break
		}

		year, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer or 'quit'.")
			continue
		}
		fmt.Println(CheckLeapYear(year))
	}

}
