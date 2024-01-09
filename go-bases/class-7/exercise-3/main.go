package main

import (
	"errors"
	"fmt"
)

func main() {
	salary := 22

	err := salaryLessThan10000(salary)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Salary is ok.")
}

func salaryLessThan10000(s int) error {
	if s < 10000 {
		return errors.New("Error: salary is less than 10000")
	}

	return nil
}
