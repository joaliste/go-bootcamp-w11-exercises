package main

import "fmt"

// Another way to define this is with a map.
var months = [12]string{
	"Enero",
	"Febrero",
	"Marzo",
	"Abril",
	"Mayo",
	"Junio",
	"Julio",
	"Agosto",
	"Septiembre",
	"Octubre",
	"Noviembre",
	"Diciembre",
}

func main() {
	fmt.Println(getMonthName(2))
}

func getMonthName(number int) string {
	if number < 1 || number > 12 {
		return "Number must be a valid month"
	}
	return months[number]
}
