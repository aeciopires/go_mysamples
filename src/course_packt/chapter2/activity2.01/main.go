package main

import (
	"errors"
	"fmt"
	"strconv"
)

func validate(number int) error {
	if number > 0 && number <= 100 {
		if number%3 == 0 && number%5 == 0 {
			return errors.New("FizzBuzz")
		} else if number%3 == 0 {
			return errors.New("Fizz")
		} else if number%5 == 0 {
			return errors.New("Buzz")
		} else {
			// Convert number to string
			result := strconv.Itoa(number)
			return errors.New(result)
		}
	}

	return errors.New("[FAIL] The number must be greater than 0 and less than 100")
}

func main() {
	for input := 1; input > 0 && input <= 100; input++ {
		fmt.Println(validate(input))
	}
}
