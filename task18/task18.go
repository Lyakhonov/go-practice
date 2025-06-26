package main

import "fmt"

func main() {
	var dayNumber int
	fmt.Print("Введите номер дня недели (1-7): ")
	fmt.Scanln(&dayNumber)

	var dayName string

	switch dayNumber {
	case 1:
		dayName = "Понедельник"
	case 2:
		dayName = "Вторник"
	case 3:
		dayName = "Среда"
	case 4:
		dayName = "Четверг"
	case 5:
		dayName = "Пятница"
	case 6:
		dayName = "Суббота"
	case 7:
		dayName = "Воскресенье"
	default:
		fmt.Printf("Ошибка: номер %d не соответствует дню недели\n", dayNumber)
		return
	}

	fmt.Printf("День недели: %s\n", dayName)
}
