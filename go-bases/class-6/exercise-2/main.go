package main

import "fmt"

type Product interface {
	Price() float64
}

type Small struct {
	cost float64
}
type Medium struct {
	cost float64
}
type Large struct {
	cost float64
}

const (
	small  = "SMALL"
	medium = "MEDIUM"
	large  = "LARGE"
)

func main() {
	smallProduct := factory("SMALL", 306)
	mediumProduct := factory("MEDIUM", 552)
	largeProduct := factory("LARGE", 706)

	fmt.Printf("Small product price: %.2f\n", smallProduct.Price())
	fmt.Printf("Medium product price: %.2f\n", mediumProduct.Price())
	fmt.Printf("Large product price: %.2f\n", largeProduct.Price())
}

func (s Small) Price() float64 {
	return s.cost
}

func (m Medium) Price() float64 {
	return m.cost * 1.03
}

func (l Large) Price() float64 {
	return l.cost*1.06 + 2500
}

func factory(productType string, cost float64) Product {
	switch productType {
	case small:
		return Small{cost: cost}
	case medium:
		return Medium{cost: cost}
	case large:
		return Large{cost: cost}
	}

	return nil
}
