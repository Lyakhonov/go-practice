package main

import "fmt"

type Student struct {
	Name     string
	Age      int
	Course   int
	AvgGrade float64
}

func NewStudent(name string, age, course int, avgGrade float64) *Student {
	return &Student{
		Name:     name,
		Age:      age,
		Course:   course,
		AvgGrade: avgGrade,
	}
}

func (s *Student) PrintInfo() {
	fmt.Printf("Студент: %s\n", s.Name)
	fmt.Printf("Возраст: %d лет\n", s.Age)
	fmt.Printf("Курс: %d\n", s.Course)
	fmt.Printf("Средний балл: %.2f\n", s.AvgGrade)
}

func (s *Student) Promote() {
	s.Course++
	fmt.Printf("%s переведен на %d курс\n", s.Name, s.Course)
}

func (s *Student) UpdateGrade(newGrade float64) {
	s.AvgGrade = newGrade
	fmt.Printf("Средний балл %s обновлен: %.2f\n", s.Name, s.AvgGrade)
}

func main() {
	student1 := NewStudent("Иван Петров", 19, 1, 4.2)
	student2 := NewStudent("Мария Сидорова", 20, 2, 4.8)

	fmt.Println("===== Информация о студентах =====")
	student1.PrintInfo()
	fmt.Println()
	student2.PrintInfo()

	fmt.Println("\n===== Перевод на следующий курс =====")
	student1.Promote()
	student2.Promote()

	// Обновляем оценки
	fmt.Println("\n===== Обновление оценок =====")
	student1.UpdateGrade(4.5)
	student2.UpdateGrade(5.0)

	// Выводим обновленную информацию
	fmt.Println("\n===== Обновленная информация =====")
	student1.PrintInfo()
	fmt.Println()
	student2.PrintInfo()
}
