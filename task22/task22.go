package main

import "fmt"

type Motor struct {
	Kind       string
	Horsepower float64
	Capacity   float64
	Active     bool
}

func (m *Motor) Activate() {
	if !m.Active {
		m.Active = true
		fmt.Printf("Мотор %s активирован\n", m.Kind)
	} else {
		fmt.Println("Мотор уже работает")
	}
}

func (m *Motor) Deactivate() {
	if m.Active {
		m.Active = false
		fmt.Printf("Мотор %s деактивирован\n", m.Kind)
	} else {
		fmt.Println("Мотор уже остановлен")
	}
}

type Vehicle struct {
	Brand  string
	Series string
	Year   int
	Power  Motor
}

func (v *Vehicle) Display() {
	fmt.Printf("Транспорт: %s %s %d года выпуска\n", v.Brand, v.Series, v.Year)
	fmt.Printf("Силовая установка: %s, %.1f л, %.1f л.с.\n",
		v.Power.Kind, v.Power.Capacity, v.Power.Horsepower)
	fmt.Printf("Статус: ")
	if v.Power.Active {
		fmt.Println("активен")
	} else {
		fmt.Println("неактивен")
	}
}

func main() {
	v6Motor := Motor{
		Kind:       "дизель",
		Horsepower: 320,
		Capacity:   3.5,
	}
	myVehicle := Vehicle{
		Brand:  "Honda",
		Series: "Accord",
		Year:   2023,
		Power:  v6Motor,
	}

	fmt.Println("Данные о транспорте:")
	myVehicle.Display()

	fmt.Println("\nАктивация мотора:")
	myVehicle.Power.Activate()
	myVehicle.Power.Activate()

	fmt.Println("\nОбновленные данные:")
	myVehicle.Display()

	fmt.Println("\nДеактивация мотора:")
	myVehicle.Power.Deactivate()
	myVehicle.Power.Deactivate()

	fmt.Println("\nИтоговая информация:")
	myVehicle.Display()
}
