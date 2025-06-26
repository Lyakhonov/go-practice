package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = math.Pi
	const e = math.E

	radius := 5.0
	circleArea := pi * math.Pow(radius, 2)
	naturalLog := math.Log(e)

	fmt.Println(pi)
	fmt.Println(e)
	fmt.Println(circleArea)
	fmt.Println(naturalLog)
}
