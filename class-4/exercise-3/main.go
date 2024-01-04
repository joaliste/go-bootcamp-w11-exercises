package main

import "fmt"

const (
	a             = "A"
	b             = "B"
	c             = "C"
	aSalary       = 3000
	bSalary       = 1500
	cSalary       = 1000
	aCompensation = 1.5
	bCompensation = 1.2
)

func main() {
	fmt.Printf("Employee salary: %.2f", getEmployeeSalary(63, c))
}

func getEmployeeSalary(m int, category string) float64 {
	hours := float64(m) / 60.0

	switch category {
	case a:
		return hours * aSalary * aCompensation
	case b:
		return hours * bSalary * bCompensation
	case c:
		return hours * cSalary
	}

	panic("Use a valid category.")
}
