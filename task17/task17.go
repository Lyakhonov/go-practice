package main

import (
	"fmt"
	"os"
)

func main() {
	var x, y float64
	var op string

	fmt.Print("Введите первое значение: ")
	fmt.Scanln(&x)
	fmt.Print("Введите операцию (+, -, ×, ÷): ")
	fmt.Scanln(&op)
	fmt.Print("Введите второе значение: ")
	fmt.Scanln(&y)

	var res float64
	var problem error

	switch op {
	case "+":
		res = x + y
	case "-":
		res = x - y
	case "×":
		res = x * y
	case "÷":
		if y == 0 {
			problem = fmt.Errorf("недопустимая операция: деление на ноль")
		} else {
			res = x / y
		}
	default:
		problem = fmt.Errorf("неподдерживаемая операция: '%s'", op)
	}

	if problem != nil {
		fmt.Println(problem)
		os.Exit(1)
	} else {
		fmt.Printf("Вычисление: %.2f %s %.2f = %.2f\n", x, op, y, res)
	}
}
