package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "И пусть никто не уйдет обиженный!"

	// 1. Подсчёт символов
	byteCount := len(text)
	fmt.Printf("Текст: \"%s\"\n", text)
	fmt.Printf("Количество символов: %d\n", byteCount)

	// 2. Поиск подстроки
	substr := "обиженный"
	index := strings.Index(text, substr)
	if index != -1 {
		fmt.Printf("Подстрока \"%s\" найдена на позиции %d\n", substr, index)
	} else {
		fmt.Printf("Подстрока \"%s\" не найдена\n", substr)
	}

	// 3. Изменение регистра
	upper := strings.ToUpper(text)
	lower := strings.ToLower(text)

	fmt.Println("\nИзменение регистра:")
	fmt.Printf("Верхний регистр: %s\n", upper)
	fmt.Printf("Нижний регистр: %s\n", lower)
}
