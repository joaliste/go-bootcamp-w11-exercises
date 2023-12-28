package main

import "fmt"

func main() {
	age := 44
	employed := true
	employedYears := 1.4
	salary := 223

	loanFeasibility(age, employed, employedYears, salary)
}

func loanFeasibility(age int, employed bool, employedYears float64, salary int) {

	if employed && age >= 22 && employedYears > 1.0 {
		switch {
		case salary > 100000:
			fmt.Println("Loan available without interest!")
		default:
			fmt.Println("Loan available with interest!")
		}
	} else {
		fmt.Println("Loan not available")
	}
}
