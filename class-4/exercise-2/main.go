package main

import "fmt"

func main() {
	fmt.Printf("Student grade point average: %.2f\n", getGPA(5.1, 6.2, 7.0, 5))
}

func getGPA(grades ...float64) float64 {
	gradesAmount := float64(len(grades))
	gradesSum := 0.0

	for _, value := range grades {
		if value < 0 {
			panic("Grades below 0 are not allowed.")
		}

		gradesSum += value
	}

	return gradesSum / gradesAmount

}
