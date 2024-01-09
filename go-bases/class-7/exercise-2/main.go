package main

import (
	"errors"
	"fmt"
)

type salaryError struct {
}

func main() {
	salary := 22

	err := salaryLessThan10000(salary)

	if err != nil {
		if errors.Is(err, &salaryError{}) {
			fmt.Println(err)
		}
		return
	}

	fmt.Println("Salary is ok.")
}

func (e *salaryError) Error() string {
	return "Error: salary is less than 10000"
}

func salaryLessThan10000(s int) error {
	if s < 10000 {
		return &salaryError{}
	}

	return nil
}
