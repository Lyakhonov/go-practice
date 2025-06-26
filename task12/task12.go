package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Исходный срез:", numbers)

	indexToRemove := 2
	numbers = removeElement(numbers, indexToRemove)
	fmt.Println("После удаления:", numbers)

	numbers = removeElement(numbers, 0)
	fmt.Println("Удаление первого элемента:", numbers)

	numbers = removeElement(numbers, len(numbers)-1)
	fmt.Println("Удаление последнего элемента:", numbers)
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		fmt.Println("Неверный индекс!")
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}
