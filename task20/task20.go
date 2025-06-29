package main

import "fmt"

func main() {
	var inputYear int
	fmt.Print("Введите интересующий год: ")
	fmt.Scanln(&inputYear)

	leap := (inputYear%4 == 0 && inputYear%100 != 0) || (inputYear%400 == 0)

	if leap {
		fmt.Printf("Год %d является високосным\n", inputYear)
	} else {
		fmt.Printf("Год %d не является високосным\n", inputYear)
	}
}
