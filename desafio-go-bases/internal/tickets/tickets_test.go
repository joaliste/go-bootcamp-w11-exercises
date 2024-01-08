package tickets

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetTotalTickets_Brazil(t *testing.T) {
	expected := 5
	country := "Brazil"
	CountryCountMap[country] = expected

	result, _ := GetTotalTickets(country)

	require.Equal(t, expected, result)
}

func TestGetTotalTickets_Santiago(t *testing.T) {
	expected := 8
	country := "Santiago"
	CountryCountMap[country] = expected

	result, _ := GetTotalTickets(country)

	require.Equal(t, expected, result)
}
