package main

import (
	"errors"
	"fmt"
)

type Customer struct {
	File        string
	Name        string
	ID          int
	PhoneNumber int
	Home        string
}

var customers = []Customer{
	{
		File:        "A.txt",
		Name:        "John",
		ID:          1,
		PhoneNumber: 555555,
		Home:        "South Street",
	},
	{
		File:        "B.txt",
		Name:        "Dina",
		ID:          2,
		PhoneNumber: 555555111,
		Home:        "North Street",
	},
}

func main() {

	defer func() {
		fmt.Println("Program finished")
		fmt.Println("Several errors were detected at runtime")
	}()

	newCustomer := Customer{
		File:        "A.txt",
		Name:        "",
		ID:          1,
		PhoneNumber: 555555,
		Home:        "South Street",
	}

	checkCustomerExistence(newCustomer)
	_, err := checkCustomerData(newCustomer)

	if err != nil {
		fmt.Println(err)
	}

}

func checkCustomerExistence(c Customer) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	for _, customer := range customers {
		if c.ID == customer.ID {
			panic("Error: client already exists")
		}
	}
}

func checkCustomerData(c Customer) (bool, error) {
	if c.Home == "" || c.Name == "" || c.File == "" || c.ID == 0 || c.PhoneNumber == 0 {
		return false, errors.New("customer have empty values")
	}

	return true, nil
}
