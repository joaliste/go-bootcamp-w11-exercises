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

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("file closed")
	}(f)

	bytes, err := io.ReadAll(f)
	if err != nil {
		panic("The file could not be read")
	}
	
	data := string(bytes)
	fmt.Println(data)
}
