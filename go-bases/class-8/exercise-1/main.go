package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	defer func() {
		fmt.Println("Execution finished")
	}()
	f, err := os.Open("customers.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := string(bytes)
	fmt.Println(data)
}
