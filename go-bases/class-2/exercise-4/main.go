package main

import "fmt"

func main() {
	var employees = map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Darío":    44,
		"Pedro":    30,
	}

	fmt.Println("La edad de Benjamin es:", employees["Benjamin"])

	over21Employees := 0
	for _, age := range employees {
		if age >= 21 {
			over21Employees++
		}
	}

	fmt.Println(over21Employees, "empleados tienen más de 21 años.")

	delete(employees, "Pedro")
	fmt.Println(employees)
}
