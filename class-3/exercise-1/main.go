package main

import "fmt"

func main() {
	fmt.Println("Salary taxes:", getSalaryTaxes(200000))
}

func getSalaryTaxes(salary float64) float64 {
	if salary > 150000 {
		return salary * 0.27
	} else if salary > 50000 {
		return salary * 0.17
	}

	return 0
}
