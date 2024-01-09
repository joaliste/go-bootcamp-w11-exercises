package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetGPA_NegativeGrade(t *testing.T) {
	require.Panics(t, func() { getGPA(1.2, 2.3, -4.1) }, "The code did not panic")
}

func TestGetGPA_getGPA(t *testing.T) {
	result := getGPA(2.4, 2.1, 5.4)
	expected := (2.4 + 2.1 + 5.4) / float64(3)
	require.Equal(t, expected, result)
}
