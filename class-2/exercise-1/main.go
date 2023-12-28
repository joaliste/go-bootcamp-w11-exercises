package main

import "fmt"

func main() {
	word := "bootcamp"
	fmt.Println("Amount of letters:", len(word))
	printLetters(word)
}

func printLetters(word string) {
	for _, letter := range word {
		fmt.Println(string(letter))
	}
}
