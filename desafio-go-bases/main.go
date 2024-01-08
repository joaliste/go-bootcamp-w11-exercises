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
		panic("File can not be read")
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
		err = tickets.ProcessRow(row)

		if err != nil {
			panic("an error occurred while processing a row")
		}
	}

	total, err := tickets.GetTotalTickets("Argentina")
	fmt.Println("Brazil total tickets:", total)
	count, err := tickets.GetCountByPeriod("9:12")
	fmt.Println("Dawn total tickets:", count)
	percentage, err := tickets.AverageDestination("Argentina")
	fmt.Printf("Brazil tickets percentage: %.2f%%\n", percentage)

}
