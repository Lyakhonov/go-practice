package main

import "fmt"

type Transport interface {
	Move(speed int)
	Stop()
}

type Car struct {
	Model string
}

func (c Car) Move(speed int) {
	fmt.Printf("Автомобиль %s движется со скоростью %d км/ч\n", c.Model, speed)
}

func (c Car) Stop() {
	fmt.Printf("Автомобиль %s остановился\n", c.Model)
}

type Bicycle struct {
	Brand string
}

func (b Bicycle) Move(speed int) {
	fmt.Printf("Велосипед %s движется со скоростью %d км/ч\n", b.Brand, speed)
}

func (b Bicycle) Stop() {
	fmt.Printf("Велосипед %s остановился\n", b.Brand)
}

type Bus struct {
	RouteNumber int
}

func (b Bus) Move(speed int) {
	fmt.Printf("Автобус маршрута %d движется со скоростью %d км/ч\n", b.RouteNumber, speed)
}

func (b Bus) Stop() {
	fmt.Printf("Автобус маршрута %d остановился\n", b.RouteNumber)
}

func testTransport(transport Transport) {
	transport.Move(60)
	transport.Stop()
	fmt.Println()
}

func main() {
	car := Car{Model: "Toyota Camry"}
	bicycle := Bicycle{Brand: "Stels"}
	bus := Bus{RouteNumber: 42}

	fmt.Println("===== Тест автомобиля =====")
	testTransport(car)

	fmt.Println("===== Тест велосипеда =====")
	testTransport(bicycle)

	fmt.Println("===== Тест автобуса =====")
	testTransport(bus)

	fmt.Println("===== Парк транспорта =====")
	transports := []Transport{car, bicycle, bus}
	for i, t := range transports {
		fmt.Printf("Транспорт %d:\n", i+1)
		t.Move(50)
		t.Stop()
		fmt.Println()
	}
}
