package main

import "fmt"

func main() {
	const max = 100
	fmt.Printf("Простые числа до %d:\n", max)
	isPrime := make([]bool, max+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false

	for i := 2; i*i <= max; i++ {
		if isPrime[i] {
			for j := i * i; j <= max; j += i {
				isPrime[j] = false
			}
		}
	}

	count := 0
	for i, prime := range isPrime {
		if prime {
			fmt.Printf("%4d", i)
			count++
			if count%10 == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Printf("\n\nВсего найдено: %d простых чисел\n", count)
}
