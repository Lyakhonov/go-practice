package main

import "fmt"

func main() {
	var year int
	fmt.Print("Введите год: ")
	fmt.Scanln(&year)

	isLeap := (year%4 == 0 && year%100 != 0) || (year%400 == 0)

	if isLeap {
		fmt.Printf("%d год - високосный\n", year)
	} else {
		fmt.Printf("%d год - не високосный\n", year)
	}
}
