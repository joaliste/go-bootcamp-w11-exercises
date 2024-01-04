package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAnimal_Dog(t *testing.T) {
	animalDog, _ := animal(dog)

	expected := dogFood * 10
	result := animalDog(10)
	require.Equal(t, expected, result)
}

func TestAnimal_Cat(t *testing.T) {
	animalCat, _ := animal(cat)

	expected := catFood * 10
	result := animalCat(10)
	require.Equal(t, expected, result)
}

func TestAnimal_Hamster(t *testing.T) {
	animalHamster, _ := animal(hamster)

	expected := hamsterFood * 10
	result := animalHamster(10)
	require.Equal(t, expected, result)
}

func TestAnimal_Spider(t *testing.T) {
	animalSpider, _ := animal(spider)

	expected := spiderFood * 10
	result := animalSpider(10)
	require.Equal(t, expected, result)
}
