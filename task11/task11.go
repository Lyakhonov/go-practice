package main

import "fmt"

func main() {
	var fruits []string

	fruits = append(fruits, "Яблоко")
	fruits = append(fruits, "Банан")
	fruits = append(fruits, "Апельсин")

	fruits = append(fruits, "Груша", "Киви")

	fmt.Println("==== Исходный срез ====")
	printSlice(fruits)

	fruits = append(fruits, "Манго", "Ананас")
	fmt.Println("\n==== После добавления новых элементов ====")
	printSlice(fruits)
}

func printSlice(slice []string) {
	for i, item := range slice {
		fmt.Printf("Индекс: %d, Значение: %s\n", i, item)
	}
	fmt.Printf("Длина: %d, Емкость: %d\n", len(slice), cap(slice))
}
