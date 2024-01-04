package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// We could generate random values for the minutes and test more cases for every category.

func TestGetEmployeeSalary_aCategory(t *testing.T) {
	expected := 3000 * 1.5 * 5.0
	require.Equal(t, expected, getEmployeeSalary(60*5, "A"))
}

func TestGetEmployeeSalary_bCategory(t *testing.T) {
	expected := 1500 * 1.2 * 5.0
	require.Equal(t, expected, getEmployeeSalary(60*5, "B"))
}

func TestGetEmployeeSalary_cCategory(t *testing.T) {
	expected := 1000 * 5.0
	require.Equal(t, expected, getEmployeeSalary(60*5, "C"))
}
