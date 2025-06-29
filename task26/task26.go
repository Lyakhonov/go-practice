package main

import "fmt"

type User struct {
	FullName string
	Years    int
}

func changeByValue(u User) {
	u.FullName = "Изменено через значение"
	u.Years = 99
	fmt.Println("\nВнутри changeByValue:")
	fmt.Printf("Адрес: %p, Содержимое: %+v\n", &u, u)
}

func changeByPointer(u *User) {
	u.FullName = "Изменено через указатель"
	u.Years = 199
	fmt.Println("\nВнутри changeByPointer:")
	fmt.Printf("Адрес: %p, Содержимое: %+v\n", u, *u)
}

func main() {
	fmt.Println("=== Примитивные типы ===")
	number := 10
	fmt.Println("\nПеред вызовом changeNumber:", number)

	changeNumber(number)
	fmt.Println("После changeNumber (передача по значению):", number)

	changeNumberPtr(&number)
	fmt.Println("После changeNumberPtr (передача по указателю):", number)

	fmt.Println("\n=== Пользовательские структуры ===")
	user := User{FullName: "Иван", Years: 30}

	fmt.Println("\nПеред changeByValue:")
	fmt.Printf("Адрес: %p, Содержимое: %+v\n", &user, user)
	changeByValue(user)
	fmt.Println("\nПосле changeByValue:")
	fmt.Printf("Адрес: %p, Содержимое: %+v\n", &user, user)

	fmt.Println("\nПеред changeByPointer:")
	fmt.Printf("Адрес: %p, Содержимое: %+v\n", &user, user)
	changeByPointer(&user)
	fmt.Println("\nПосле changeByPointer:")
	fmt.Printf("Адрес: %p, Содержимое: %+v\n", &user, user)
}

func changeNumber(n int) {
	n = 100
}

func changeNumberPtr(n *int) {
	*n = 200
}
