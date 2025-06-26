package main

import (
	"fmt"
	"sort"
)

func FindElement(slice []int, value int) int {
	for index, element := range slice {
		if element == value {
			return index
		}
	}
	return -1
}

func SortSlice(slice []int) []int {
	sorted := make([]int, len(slice))
	copy(sorted, slice)
	sort.Ints(sorted)
	return sorted
}

func FilterSlice(slice []int, condition func(int) bool) []int {
	var filtered []int
	for _, element := range slice {
		if condition(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered
}

func main() {
	numbers := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}

	fmt.Println("Исходный срез:", numbers)
	searchValue := 8
	index := FindElement(numbers, searchValue)
	if index != -1 {
		fmt.Printf("Элемент %d найден на позиции %d\n", searchValue, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", searchValue)
	}

	sorted := SortSlice(numbers)
	fmt.Println("Отсортированный срез:", sorted)

	evenNumbers := FilterSlice(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Четные числа:", evenNumbers)

	greaterThanFive := FilterSlice(numbers, func(n int) bool {
		return n > 5
	})
	fmt.Println("Числа больше 5:", greaterThanFive)
}
