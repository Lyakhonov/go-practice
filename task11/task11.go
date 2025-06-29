package main

import "fmt"

func main() {
	var items []string

	items = append(items, "Яблоко")
	items = append(items, "Банан")
	items = append(items, "Апельсин")

	items = append(items, "Груша", "Киви")

	fmt.Println("=== Начальный список ===")
	showSlice(items)

	items = append(items, "Манго", "Ананас")
	fmt.Println("\n=== После добавления элементов ===")
	showSlice(items)
}

func showSlice(s []string) {
	for idx, val := range s {
		fmt.Printf("Индекс: %d, Значение: %s\n", idx, val)
	}
	fmt.Printf("Длина среза: %d, Вместимость: %d\n", len(s), cap(s))
}
