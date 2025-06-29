package main

import "fmt"

func main() {
	scores := make(map[string]int)

	addScore(scores, "Алексей", 5)
	addScore(scores, "Мария", 4)
	addScore(scores, "Иван", 3)
	addScore(scores, "Алексей", 4)

	searchScore(scores, "Мария")
	searchScore(scores, "Петр")

	deleteScore(scores, "Иван")
	deleteScore(scores, "Ольга")

	fmt.Println("\nТекущие оценки студентов:")
	for student, score := range scores {
		fmt.Printf("%s: %d\n", student, score)
	}
}

func addScore(m map[string]int, name string, grade int) {
	m[name] = grade
	fmt.Printf("Оценка для %s установлена на %d\n", name, grade)
}

func searchScore(m map[string]int, name string) {
	if grade, ok := m[name]; ok {
		fmt.Printf("Студент %s имеет оценку %d\n", name, grade)
	} else {
		fmt.Printf("Студент %s не найден\n", name)
	}
}

func deleteScore(m map[string]int, name string) {
	if _, ok := m[name]; ok {
		delete(m, name)
		fmt.Printf("Оценка студента %s удалена\n", name)
	} else {
		fmt.Printf("Студент %s отсутствует, удаление невозможно\n", name)
	}
}
