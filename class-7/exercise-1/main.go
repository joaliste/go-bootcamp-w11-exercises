package main

import "fmt"

type salaryError struct {
}

func main() {
	salary := 12033993
	if salary < 150000 {
		err := &salaryError{}
		fmt.Println(err)
		return
	}

	fmt.Println("Must pay tax")
}

func (e *salaryError) Error() string {
	return "Error: the salary entered does not reach the taxable minimum"
}
