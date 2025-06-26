package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func printArea(shape Shape) {
	fmt.Printf("Площадь фигуры: %.2f\n", shape.Area())
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	fmt.Println("Прямоугольник:")
	fmt.Printf("Ширина: %.2f, Высота: %.2f\n", rect.Width, rect.Height)
	printArea(rect)

	fmt.Println("\nКруг:")
	fmt.Printf("Радиус: %.2f\n", circle.Radius)
	printArea(circle)

	fmt.Println("\nФормы:")
	shapes := []Shape{rect, circle}
	for i, shape := range shapes {
		fmt.Printf("Фигура %d: площадь = %.2f\n", i+1, shape.Area())
	}
}
