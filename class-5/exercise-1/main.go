package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{
	{ID: 1, Name: "p1", Price: 2.3, Description: "p1_desc", Category: "A"},
	{ID: 2, Name: "p2", Price: 5.0, Description: "p2_desc", Category: "B"},
}

func main() {
	product := Product{
		ID:          3,
		Name:        "p3",
		Price:       221.3,
		Description: "p3_desc",
		Category:    "A",
	}

	product.Save()
	product.GetAll()

	element, err := getById(2)

	if err != nil {
		return
	}

	fmt.Printf("%v\n", element)
}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	fmt.Printf("%v\n", Products)
}

func getById(id int) (Product, error) {
	for _, element := range Products {
		if element.ID == id {
			return element, nil
		}
	}

	return Product{}, errors.New("product not found")
}
