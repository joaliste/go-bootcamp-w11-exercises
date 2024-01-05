package main

import (
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
	if s < 150000 {
		return fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %d\n", s)
	}

	return nil
}
