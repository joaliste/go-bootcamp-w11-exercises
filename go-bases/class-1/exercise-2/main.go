package main

import "fmt"

func main() {
	var temperature float64
	var humidity float64
	var pressure float64

	temperature = 23.5
	humidity = 3.241
	pressure = 123.4

	fmt.Printf(
		"Temperatura: %.2f, humedad: %.2f, presión: %.2f \n",
		temperature,
		humidity,
		pressure,
	)
}
