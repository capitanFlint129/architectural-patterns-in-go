package main

import (
	"fmt"

	"architectural-patterns-in-go/pkg/builder"
)

func main() {
	// строим машины
	carBuilder := builder.NewCarBuilder()
	director := builder.NewCarDirector()
	director.SetBuilder(carBuilder)
	director.ConstructSuvCar()
	suvCar := carBuilder.GetResult()
	director.ConstructSportsCar()
	sportsCar := carBuilder.GetResult()

	// пишем инструкции
	manualBuilder := builder.NewManualBuilder()
	director.SetBuilder(manualBuilder)
	director.ConstructSuvCar()
	suvCarManual := manualBuilder.GetResult()
	director.ConstructSportsCar()
	sportsCarManual := manualBuilder.GetResult()

	fmt.Println(suvCar)
	fmt.Println(sportsCar)
	fmt.Println(suvCarManual)
	fmt.Println(sportsCarManual)
}
