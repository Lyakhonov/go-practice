package main

import "fmt"

type Transport interface {
	Drive(speed int)
	Halt()
}

type Car struct {
	Name string
}

func (car Car) Drive(speed int) {
	fmt.Printf("Машина %s едет со скоростью %d км/ч\n", car.Name, speed)
}

func (car Car) Halt() {
	fmt.Printf("Машина %s остановилась\n", car.Name)
}

type Bicycle struct {
	Make string
}

func (bicycle Bicycle) Drive(speed int) {
	fmt.Printf("Велосипед марки %s движется со скоростью %d км/ч\n", bicycle.Make, speed)
}

func (bicycle Bicycle) Halt() {
	fmt.Printf("Велосипед %s остановился\n", bicycle.Make)
}

type Bus struct {
	Line int
}

func (bus Bus) Drive(speed int) {
	fmt.Printf("Автобус линии %d едет со скоростью %d км/ч\n", bus.Line, speed)
}

func (bus Bus) Halt() {
	fmt.Printf("Автобус линии %d остановился\n", bus.Line)
}

func testVehicle(t Transport) {
	t.Drive(60)
	t.Halt()
	fmt.Println()
}

func main() {
	myCar := Car{Name: "Toyota Camry"}
	myBike := Bicycle{Make: "Stels"}
	myBus := Bus{Line: 42}

	fmt.Println("=== Тестируем автомобиль ===")
	testVehicle(myCar)

	fmt.Println("=== Тестируем велосипед ===")
	testVehicle(myBike)

	fmt.Println("=== Тестируем автобус ===")
	testVehicle(myBus)

	fmt.Println("=== Все транспортные средства ===")
	vehicles := []Transport{myCar, myBike, myBus}
	for idx, vehicle := range vehicles {
		fmt.Printf("Транспортное средство №%d:\n", idx+1)
		vehicle.Drive(50)
		vehicle.Halt()
		fmt.Println()
	}
}
