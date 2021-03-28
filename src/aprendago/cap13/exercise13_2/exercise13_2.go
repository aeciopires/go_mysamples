package	main

import (
	"fmt"
)

//func (receiver) identifier(parameters) (returns) { code }

func sumInteger(numbers ...int) int{
	total := 0
	fmt.Print("Numbers: ")
	for aux, number := range numbers {
		fmt.Print(number)
		total += number
		if aux < (len(numbers)-1){
			fmt.Print(" + ")
		}
	}
	fmt.Println("")
	return total
}

func sumSliceInterger(numbers []int) int{
	total := 0
	printSlice(numbers)
	for _, number := range numbers {
		total += number
	}
	return total
}

func printSlice(numbers []int) {
	fmt.Printf("length=%d, capacity=%d, slice: %v\n", len(numbers), cap(numbers), numbers)
}

func main(){
	nums := []int{5, 4, 2, 4, 10, 15}
	fmt.Println("Sum of integers: ",sumInteger(nums...))
	fmt.Println("Sum of integers in the slice: ",sumSliceInterger(nums))
}