package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person
}

func main() {
	person := Person{
		ID:          1,
		Name:        "John",
		DateOfBirth: "20 of july of 1923",
	}

	employee := Employee{
		ID:       1,
		Position: "Door",
		Person:   person,
	}

	employee.PrintEmployee()

	// Using composition
	employee.PrintPerson()
}

func (e Employee) PrintEmployee() {
	fmt.Printf("%v\n", e)
}

func (p Person) PrintPerson() {
	fmt.Printf("%v\n", p)
}
