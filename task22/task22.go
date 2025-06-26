package main

import "fmt"

type Engine struct {
	Type      string
	Power     float64
	Volume    float64
	IsRunning bool
}

func (e *Engine) Start() {
	if !e.IsRunning {
		e.IsRunning = true
		fmt.Printf("Двигатель %s запущен\n", e.Type)
	} else {
		fmt.Println("Двигатель уже работает")
	}
}

func (e *Engine) Stop() {
	if e.IsRunning {
		e.IsRunning = false
		fmt.Printf("Двигатель %s остановлен\n", e.Type)
	} else {
		fmt.Println("Двигатель уже остановлен")
	}
}

type Car struct {
	Make       string
	Model      string
	Year       int
	EngineInfo Engine
}

func (c *Car) PrintInfo() {
	fmt.Printf("Автомобиль: %s %s %d года\n", c.Make, c.Model, c.Year)
	fmt.Printf("Двигатель: %s, %.1f л, %.1f л.с.\n",
		c.EngineInfo.Type, c.EngineInfo.Volume, c.EngineInfo.Power)
	fmt.Printf("Состояние: ")
	if c.EngineInfo.IsRunning {
		fmt.Println("заведен")
	} else {
		fmt.Println("не заведен")
	}
}

func main() {
	v8Engine := Engine{
		Type:   "бензин",
		Power:  420,
		Volume: 5.0,
	}
	myCar := Car{
		Make:       "Toyota",
		Model:      "Camry",
		Year:       2022,
		EngineInfo: v8Engine,
	}

	fmt.Println("Информация об автомобиле:")
	myCar.PrintInfo()

	fmt.Println("\nЗапуск двигателя:")
	myCar.EngineInfo.Start()
	myCar.EngineInfo.Start()

	fmt.Println("\nОбновленная информация:")
	myCar.PrintInfo()

	fmt.Println("\nОстановка двигателя:")
	myCar.EngineInfo.Stop()
	myCar.EngineInfo.Stop()

	fmt.Println("\nФинальная информация:")
	myCar.PrintInfo()
}
