package main

import (
	"fmt"
)

func main() {
	limit := 0
	year := 2019
	months30days := []int{4, 6, 9, 11}
	months31days := []int{1, 3, 5, 7, 8, 10, 12}

	for months := 1; months <= 12; months++ {
		if (months == 2) && (multipleFour(year)) {
			limit = 29
			fmt.Println("####### Ano Bissexto")
		} else if months == 2 {
			limit = 28
			fmt.Println("####### Ano comum")
		} else if containsMonth(months30days, months) {
			limit = 30
		} else if containsMonth(months31days, months) {
			limit = 31
		}

		fmt.Println("------ BEGIN-MONTH")
		fmt.Println("Mês: ", months)

		for days := 1; days <= limit; days++ {
			fmt.Println("------ BEGIN-DAY")
			fmt.Println("Dia: ", days)
			for hours := 0; hours < 24; hours++ {
				fmt.Println("Hora: ", hours)
				fmt.Println("Minutos: ")
				for min := 0; min < 60; min++ {
					fmt.Print(" ", min)
				}
				fmt.Println(" ")
			}
			fmt.Println(" ")
			fmt.Println("--- END-DAY ---")
		}
		fmt.Println(" ")
		fmt.Println("------ END-MONTH ------")
	}
}

func containsMonth(myarray []int, num int) bool {
	for _, number := range myarray {
		if num == number {
			return true
		}
	}
	return false
}

func multipleFour(number int) bool {
	result := number % 4
	if result == 0 {
		return true
	}
	return false
}
