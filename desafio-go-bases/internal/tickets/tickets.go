package tickets

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// CountryCountMap map with the countries count
var CountryCountMap = make(map[string]int)
var totalTickets = 0

// constants for the allowed periods.
const (
	dawn      = "DAWN"
	morning   = "MORNING"
	afternoon = "AFTERNOON"
	night     = "NIGHT"
)

// PeriodCountMap map with the periods count
var PeriodCountMap = map[string]int{
	dawn:      0,
	morning:   0,
	afternoon: 0,
	night:     0,
}

// GetTotalTickets this function get the count of the tickets sold for a country from the map CountryCountMap.
func GetTotalTickets(country string) (int, error) {
	count, ok := CountryCountMap[country]
	if !ok {
		return 0, errors.New("country does not exist")
	}

	return count, nil
}

// GetCountByPeriod return the amount of tickets sold for a specific period of time.
func GetCountByPeriod(time string) (int, error) {
	period, err := getPeriod(time)
	if err != nil {
		return 0, err
	}

	count, ok := PeriodCountMap[period]
	if !ok {
		return 0, errors.New("period does not exist")
	}

	return count, nil
}

// AverageDestination returns the percentage of flights that belongs to a specific country (destination).
func AverageDestination(destination string) (float64, error) {
	count, err := GetTotalTickets(destination)

	if err != nil {
		return 0, errors.New("destination does not exist")
	}

	return float64(count) / float64(totalTickets) * 100, nil
}

// ProcessRow process the information of csv's row. It updates the countries and period counts.
func ProcessRow(row []string) error {
	updateCountries(row[3])
	err := updatePeriodMap(row[4])
	if err != nil {
		return err
	}

	totalTickets++

	return nil
}

// updateCountries update the country map and increase the count of the country
func updateCountries(country string) {
	_, ok := CountryCountMap[country]
	if ok {
		CountryCountMap[country] += 1
	} else {
		CountryCountMap[country] = 1
	}
}

// updatePeriodMap update the period map and increase the count of the period
func updatePeriodMap(time string) error {
	time, err := getPeriod(time)
	if err != nil {
		return err
	}
	_, ok := PeriodCountMap[time]

	if !ok {
		return errors.New("time period not found in map")
	}
	PeriodCountMap[time] += 1

	return nil
}

// getPeriod returns a string with the period of a time specified in a string with format "H:mm"
func getPeriod(time string) (string, error) {
	hour := strings.Split(time, ":")[0]
	var period string

	intHour, err := strconv.Atoi(hour)
	if err != nil {
		return "", err
	}

	if intHour >= 0 && intHour <= 6 {
		period = dawn
	} else if intHour >= 7 && intHour <= 12 {
		period = morning
	} else if intHour >= 13 && intHour <= 19 {
		period = afternoon
	} else if intHour >= 20 && intHour <= 23 {
		period = night
	} else {
		err := fmt.Errorf("invalid hour value: %d", intHour)
		return "", err
	}

	return period, nil
}
