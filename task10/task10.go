package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Счастье для всех, даром, и пусть никто не уйдет обиженный!"

	words := strings.Fields(sentence)

	fmt.Printf("Предложение: \"%s\"\n", sentence)
	fmt.Println("\nСлова:")

	for i, word := range words {
		fmt.Printf("%d. %s\n", i+1, word)
	}
}
