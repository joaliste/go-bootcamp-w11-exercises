package main

import "fmt"

const (
	dog     = "dog"
	cat     = "cat"
	spider  = "spider"
	hamster = "hamster"
)

const (
	dogFood     = 10.0
	catFood     = 5.0
	spiderFood  = 0.15
	hamsterFood = 0.25
)

func main() {
	animalDog, msg := animal(dog)
	if msg != "" {
		panic(msg)
	}
	animalCat, msg := animal(cat)
	if msg != "" {
		panic(msg)
	}

	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)

	fmt.Printf("Amount of food: %.2fkg.", amount)
}

func animal(a string) (func(value float64) float64, string) {
	var foodAmount float64
	switch a {
	case dog:
		foodAmount = dogFood
	case spider:
		foodAmount = spiderFood
	case cat:
		foodAmount = catFood
	case hamster:
		foodAmount = hamsterFood
	}

	if foodAmount == 0.0 {
		return nil, "Animal not allowed."
	}

	return func(value float64) float64 {
		return foodAmount * value
	}, ""
}
