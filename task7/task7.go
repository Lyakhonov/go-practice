package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var i int = -42
	var i8 int8 = -128
	var ui uint = 42
	var ui8 uint8 = 255
	var r rune = '⌘'
	var b byte = 65
	var f32 float32 = 3.141592653589793
	var f64 float64 = 3.141592653589793
	var c64 complex64 = complex(1, 2)
	var c128 complex128 = cmplx.Sqrt(-5 + 12i)
	var s string = "Привет"
	var flag bool = true

	fmt.Println("===== Целочисленные типы =====")
	fmt.Printf("int:    %d\n", i)
	fmt.Printf("int8:   %d\n", i8)
	fmt.Printf("uint:   %d\n", ui)
	fmt.Printf("uint8:  %d\n", ui8)
	fmt.Printf("rune:   %U '%c'\n", r, r)
	fmt.Printf("byte:   %d '%c'\n", b, b)

	fmt.Println("\n===== Числа с плавающей точкой =====")
	fmt.Printf("float32: %.7f\n", f32)
	fmt.Printf("float64: %.15f\n", f64)

	fmt.Println("\n===== Комплексные числа =====")
	fmt.Printf("complex64:  %v\n", c64)
	fmt.Printf("complex128: %v\n", c128)

	fmt.Println("\n===== Строки и логические значения =====")
	fmt.Printf("string: %s\n", s)
	fmt.Printf("bool:   %t\n", flag)
}
