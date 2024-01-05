package tickets

import (
	"errors"
	"strconv"
	"strings"
)

type Ticket struct {
}

var CountryCountMap = make(map[string]int)
var totalTickets = 0

const (
	dawn      = "DAWN"
	morning   = "MORNING"
	afternoon = "AFTERNOON"
	night     = "NIGHT"
)

var PeriodCountMap = map[string]int{
	dawn:      0,
	morning:   0,
	afternoon: 0,
	night:     0,
}

// ejemplo 1
func GetTotalTickets(country string) (int, error) {
	count, ok := CountryCountMap[country]
	if !ok {
		return 0, errors.New("country does not exist")
	}

	return count, nil
}

// ejemplo 2
func GetCountByPeriod(time string) (int, error) {
	period := getPeriod(time)
	count, ok := PeriodCountMap[period]
	if !ok {
		return 0, errors.New("period does not exist")
	}

	return count, nil
}

// ejemplo 3
func AverageDestination(destination string) (float64, error) {
	count, err := GetTotalTickets(destination)

	if err != nil {
		return 0, errors.New("destination does not exist")
	}

	return float64(count) / float64(totalTickets) * 100, nil
}

func ProcessRow(row []string) {
	updateCountries(row[3])
	updatePeriodMap(row[4])
	totalTickets++
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

func updatePeriodMap(time string) {
	time = getPeriod(time)
	_, ok := PeriodCountMap[time]

	if !ok {
		panic("invalid time")
	}
	PeriodCountMap[time] += 1
}

func getPeriod(time string) string {
	hour := strings.Split(time, ":")[0]
	var period string

	intHour, err := strconv.Atoi(hour)
	if err != nil {
		panic(err)
	}

	if intHour >= 0 && intHour <= 6 {
		period = dawn
	} else if intHour >= 7 && intHour <= 12 {
		period = morning
	} else if intHour >= 13 && intHour <= 19 {
		period = afternoon
	} else if intHour >= 20 && intHour <= 23 {
		period = night
	}

	return period
}
