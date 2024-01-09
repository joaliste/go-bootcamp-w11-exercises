package main

import "fmt"

type Student struct {
	Name    string
	Surname string
	DNI     string
	Date    string
}

func main() {
	student := Student{
		Name:    "John",
		Surname: "Perez",
		DNI:     "123456-1",
		Date:    "12-02-21",
	}
	student.detail()
}

func (s Student) detail() {
	fmt.Printf("%+v\n", s)
}
