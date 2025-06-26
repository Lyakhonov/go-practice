package main

import "fmt"

func main() {
	grades := make(map[string]int)

	addGrade(grades, "Алексей", 5)
	addGrade(grades, "Мария", 4)
	addGrade(grades, "Иван", 3)
	addGrade(grades, "Алексей", 4)

	findGrade(grades, "Мария")
	findGrade(grades, "Петр")

	removeGrade(grades, "Иван")
	removeGrade(grades, "Ольга")

	fmt.Println("\nТекущие оценки:")
	for student, grade := range grades {
		fmt.Printf("%s: %d\n", student, grade)
	}
}

func addGrade(grades map[string]int, name string, grade int) {
	grades[name] = grade
	fmt.Printf("Добавлена оценка для %s: %d\n", name, grade)
}

func findGrade(grades map[string]int, name string) {
	if grade, exists := grades[name]; exists {
		fmt.Printf("Студент %s имеет оценку %d\n", name, grade)
	} else {
		fmt.Printf("Студент %s не найден\n", name)
	}
}

func removeGrade(grades map[string]int, name string) {
	if _, exists := grades[name]; exists {
		delete(grades, name)
		fmt.Printf("Оценка студента %s удалена\n", name)
	} else {
		fmt.Printf("Студент %s не найден, удаление невозможно\n", name)
	}
}
