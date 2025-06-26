package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x := 10
	y := 20

	fmt.Println("До обмена:")
	fmt.Printf("x = %d, y = %d\n", x, y)

	swap(&x, &y)

	fmt.Println("\nПосле обмена:")
	fmt.Printf("x = %d, y = %d\n", x, y)

}
