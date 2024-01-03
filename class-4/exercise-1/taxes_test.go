package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetSalaryTaxes_SalaryUnder50k(t *testing.T) {
	expected := 0.0
	result := getSalaryTaxes(20)

	// Without testify
	//if expected != result {
	//	t.Fatal("Taxes of a salary below 50k must be 0.")
	//}

	require.Equal(t, expected, result)
}

func TestGetSalaryTaxes_SalaryOver50k(t *testing.T) {
	expected := 55000 * 0.17
	result := getSalaryTaxes(55000)

	// Without testify
	//if expected != result {
	//	t.Fatal("Taxes of a salary over 50k must be the salary multiplied by 0.17.")
	//}

	require.Equal(t, expected, result)
}

func TestGetSalaryTaxes_SalaryOver150k(t *testing.T) {
	expected := 160000 * 0.27
	result := getSalaryTaxes(160000)

	// Without testify
	//if expected != result {
	//	t.Fatal("Taxes of a salary over 50k must be the salary multiplied by 0.27.")
	//}

	require.Equal(t, expected, result)
}
