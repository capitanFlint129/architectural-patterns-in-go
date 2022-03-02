package main

import (
	"fmt"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/builder"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/product"
)

var config = builder.Config{
	SuvCarParameters: builder.CarParameters{
		SeatsNumber:       6,
		EnginePower:       1000,
		TripComputerModel: "SUV car computer",
		GpsModel:          "Super GPS",
	},
	SportCarParameters: builder.CarParameters{
		SeatsNumber:       1,
		EnginePower:       700,
		TripComputerModel: "Super sports car computer",
		GpsModel:          "Cheap GPS",
	},
}

func main() {
	// строим машины
	carBuilder := builder.NewCarBuilder(product.NewCar)
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
	manualBuilder := builder.NewManualBuilder(product.NewManual)
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
