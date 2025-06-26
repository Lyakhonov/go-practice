package main

import "fmt"

func FindLongestString(strings ...string) string {
	if len(strings) == 0 {
		return ""
	}

	longest := strings[0]
	for _, s := range strings[1:] {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}

func main() {
	result1 := FindLongestString("apple", "banana", "cherry", "date")
	fmt.Println("Самая длинная строка:", result1)

	fruits := []string{"яблоко", "апельсин", "клубника", "персик"}
	result2 := FindLongestString(fruits...)
	fmt.Println("Самая длинная строка:", result2)

}
