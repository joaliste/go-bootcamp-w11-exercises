package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(getSalary(100, 300))
}

func getSalary(hours int, hourCost int) (float64, error) {
	if hours < 80 {
		return 0, errors.New("Error: the worker cannot have worked less than 80 hours per month")
	}

	salary := float64(hours * hourCost)

	if salary < 150000 {
		return 0, fmt.Errorf(
			"Error: the minimum taxable amount is 150,000 and the salary entered is: %.2f\n",
			salary,
		)
	}

	salary = salary * 0.9
	return salary, nil
}
