package main

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	minFunc, _ := operation(minimum)
	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("Min value: %.2f\n", minValue)

	maxFunc, _ := operation(maximum)
	maxValue := maxFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("Max value: %.2f\n", maxValue)

	avgFunc, _ := operation(average)
	avgValue := avgFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("Avg value: %.2f\n", avgValue)
}

func getMin(values ...float64) float64 {
	currentMin := values[0]

	for _, value := range values {
		if value < currentMin {
			currentMin = value
		}
	}

	return currentMin
}

func getAverage(values ...float64) float64 {
	valuesLen := float64(len(values))
	valuesSum := 0.0

	for _, value := range values {
		valuesSum += value
	}

	return valuesSum / valuesLen
}

func getMax(values ...float64) float64 {
	currentMax := values[0]

	for _, value := range values {
		if value > currentMax {
			currentMax = value
		}
	}

	return currentMax
}

func operation(opType string) (func(values ...float64) float64, string) {
	switch opType {
	case minimum:
		return getMin, ""
	case maximum:
		return getMax, ""
	case average:
		return getAverage, ""
	}

	return nil, "Operation not defined"
}
