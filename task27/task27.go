package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthDate time.Time
	Course    int
	AvgGrade  float64
}

func NewStudent(name string, birthYear, birthMonth, birthDay, course int, avgGrade float64) *Student {
	birthDate := time.Date(birthYear, time.Month(birthMonth), birthDay, 0, 0, 0, 0, time.UTC)
	return &Student{
		Name:      name,
		BirthDate: birthDate,
		Course:    course,
		AvgGrade:  avgGrade,
	}
}

func (s *Student) CalculateAge() int {
	now := time.Now()
	years := now.Year() - s.BirthDate.Year()

	birthdayThisYear := time.Date(now.Year(), s.BirthDate.Month(), s.BirthDate.Day(), 0, 0, 0, 0, time.UTC)
	if now.Before(birthdayThisYear) {
		years--
	}
	return years
}

func (s *Student) GetStatus() string {
	switch {
	case s.AvgGrade >= 4.5:
		return "отличник"
	case s.AvgGrade >= 3.5:
		return "хорошист"
	case s.AvgGrade >= 3.0:
		return "троечник"
	default:
		return "двоечник"
	}
}

func (s *Student) PrintInfo() {
	fmt.Printf("\nСтудент: %s\n", s.Name)
	fmt.Printf("Дата рождения: %s\n", s.BirthDate.Format("2006-01-02"))
	fmt.Printf("Возраст: %d лет\n", s.CalculateAge())
	fmt.Printf("Курс: %d\n", s.Course)
	fmt.Printf("Средний балл: %.2f\n", s.AvgGrade)
	fmt.Printf("Статус: %s\n", s.GetStatus())
}

func main() {
	student1 := NewStudent("Иван Петров", 2003, 5, 15, 2, 4.8)
	student2 := NewStudent("Мария Сидорова", 2005, 11, 3, 1, 3.7)
	student3 := NewStudent("Алексей Иванов", 2004, 2, 28, 3, 2.9)

	fmt.Println("Информация о студентах")
	student1.PrintInfo()
	student2.PrintInfo()
	student3.PrintInfo()
}
