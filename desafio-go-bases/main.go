package main

import (
	"encoding/csv"
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("tickets.csv")

	if err != nil {
		fmt.Errorf("Error while reading the file", err)
		return
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		tickets.ProcessRow(row)
	}

	total, err := tickets.GetTotalTickets("Brazil")
	fmt.Println("Brazil total tickets:", total)
	count, err := tickets.GetCountByPeriod("9:12")
	fmt.Println("Dawn total tickets:", count)
	percentage, err := tickets.AverageDestination("Brazil")
	fmt.Printf("Brazil tickets percentage: %.2f%%\n", percentage)

}
