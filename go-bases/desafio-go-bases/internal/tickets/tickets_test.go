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

func TestGetTotalTickets_CountryNotFound(t *testing.T) {
	country := "Italy"

	_, err := GetTotalTickets(country)

	require.Errorf(t, err, "country does not exist")
}

func TestGetCountByPeriod_dawnCount(t *testing.T) {
	periodTime := "2:32"
	expected := 5
	PeriodCountMap[dawn] = expected
	result, _ := GetCountByPeriod(periodTime)

	require.Equal(t, expected, result)
}

func TestGetPeriod_dawnPeriod(t *testing.T) {
	periodTime := "2:32"
	expected := dawn
	result, _ := getPeriod(periodTime)

	require.Equal(t, expected, result)
}

func TestGetPeriod_invalidTime(t *testing.T) {
	periodTime := "2222"
	_, err := getPeriod(periodTime)

	require.Errorf(t, err, "invalid hour format: 2222")
}

func TestAverageDestination_brazilAverageDestination(t *testing.T) {
	firstCountry := "Brazil"
	secondCountry := "Santiago"
	firstCountryCount := 10
	secondCountryCount := 5
	totalTickets = firstCountryCount + secondCountryCount

	CountryCountMap[firstCountry] = firstCountryCount
	CountryCountMap[secondCountry] = secondCountryCount

	expected := float64(firstCountryCount) / float64(totalTickets) * 100
	result, _ := AverageDestination("Brazil")
	require.Equal(t, expected, result)
}
