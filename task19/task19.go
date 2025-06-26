package main

import (
	"fmt"
)

func main() {
	var a, b float64 = 15, 20

	fmt.Printf("\n===== Арифметические операции =====\n")
	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
	fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)

	if b != 0 {
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
	} else {
		fmt.Printf("%.2f / %.2f = ошибка (деление на ноль)\n", a, b)
	}
}
