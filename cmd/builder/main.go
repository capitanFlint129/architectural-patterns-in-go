package main

import (
	"fmt"

	"architectural-patterns-in-go/pkg/builder"
	"architectural-patterns-in-go/pkg/product"
)

var config = Config{
	SuvCarParameters: CarParameters{
		6,
		1000,
		"SUV car computer",
		"Super GPS",
	},
	SportCarParameters: CarParameters{
		1,
		700,
		"Super sports car computer",
		"Cheap GPS",
	},
}

func main() {
	// строим машины
	carBuilder := builder.NewCarBuilder(product.CarCreator)
	director := builder.NewCarDirector()
	director.SetBuilder(carBuilder)
	director.ConstructSuvCar(
		config.SuvCarParameters.SeatsNumber,
		config.SuvCarParameters.EnginePower,
		config.SuvCarParameters.TripComputerModel,
		config.SuvCarParameters.GpsModel,
	)
	suvCar := carBuilder.GetResult()
	director.ConstructSportsCar(
		config.SportCarParameters.SeatsNumber,
		config.SportCarParameters.EnginePower,
		config.SportCarParameters.TripComputerModel,
		config.SportCarParameters.GpsModel,
	)
	sportsCar := carBuilder.GetResult()

	// пишем инструкции
	manualBuilder := builder.NewManualBuilder(product.ManualCreator)
	director.SetBuilder(manualBuilder)
	director.ConstructSuvCar(
		config.SuvCarParameters.SeatsNumber,
		config.SuvCarParameters.EnginePower,
		config.SuvCarParameters.TripComputerModel,
		config.SuvCarParameters.GpsModel,
	)
	suvCarManual := manualBuilder.GetResult()
	director.ConstructSportsCar(
		config.SportCarParameters.SeatsNumber,
		config.SportCarParameters.EnginePower,
		config.SportCarParameters.TripComputerModel,
		config.SportCarParameters.GpsModel,
	)
	sportsCarManual := manualBuilder.GetResult()

	fmt.Println(suvCar)
	fmt.Println(sportsCar)
	fmt.Println(suvCarManual)
	fmt.Println(sportsCarManual)
}
