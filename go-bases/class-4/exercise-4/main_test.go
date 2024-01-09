package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOperation_minimum(t *testing.T) {
	values := []float64{3, 2, 3, 3, 4, 10, 2, 4, 5}
	op := "minimum"
	expected := 2.0

	minFunc, _ := operation(op)
	result := minFunc(values...)

	require.Equal(t, expected, result)
}

func TestOperation_average(t *testing.T) {
	values := []float64{2, 3, 3, 4, 10, 2, 4, 5}
	op := "average"
	expected := (2.0 + 3.0 + 3.0 + 4.0 + 10.0 + 2.0 + 4.0 + 5.0) / 8.0

	minFunc, _ := operation(op)
	result := minFunc(values...)

	require.Equal(t, expected, result)
}

func TestOperation_max(t *testing.T) {
	values := []float64{2, 3, 3, 4, 10, 2, 4, 5}
	op := "maximum"
	expected := 10.0

	minFunc, _ := operation(op)
	result := minFunc(values...)

	require.Equal(t, expected, result)
}
