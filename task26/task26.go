package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func modifyValue(p Person) {
	p.Name = "Изменено по значению"
	p.Age = 100
	fmt.Println("\nВнутри modifyValue:")
	fmt.Printf("Адрес: %p, Значение: %+v\n", &p, p)
}

func modifyPointer(p *Person) {
	p.Name = "Изменено по указателю"
	p.Age = 200
	fmt.Println("\nВнутри modifyPointer:")
	fmt.Printf("Адрес: %p, Значение: %+v\n", p, *p)
}

func main() {
	fmt.Println("==== Базовые типы ====")
	val := 10
	fmt.Println("\nДо modifyInt:", val)

	modifyInt(val)
	fmt.Println("После modifyInt (передача по значению):", val)

	modifyIntPtr(&val)
	fmt.Println("После modifyIntPtr (передача по указателю):", val)

	fmt.Println("\n==== Структуры ====")
	person := Person{Name: "Иван", Age: 30}

	fmt.Println("\nДо modifyValue:")
	fmt.Printf("Адрес: %p, Значение: %+v\n", &person, person)
	modifyValue(person)
	fmt.Println("\nПосле modifyValue:")
	fmt.Printf("Адрес: %p, Значение: %+v\n", &person, person)

	fmt.Println("\nДо modifyPointer:")
	fmt.Printf("Адрес: %p, Значение: %+v\n", &person, person)
	modifyPointer(&person)
	fmt.Println("\nПосле modifyPointer:")
	fmt.Printf("Адрес: %p, Значение: %+v\n", &person, person)
}

func modifyInt(n int) {
	n = 100
}

func modifyIntPtr(n *int) {
	*n = 200
}
