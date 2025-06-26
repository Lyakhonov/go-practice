package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Go это язык программирования. Go разработан в Google. Язык Go прост и эффективен."

	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, "!", "")
	text = strings.ReplaceAll(text, "?", "")

	words := strings.Fields(text)

	wordFrequency := make(map[string]int)

	for _, word := range words {
		wordFrequency[word]++
	}

	fmt.Println("Частота встречаемости слов:")
	for word, count := range wordFrequency {
		fmt.Printf("%s: %d\n", word, count)
	}
}
