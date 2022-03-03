package main

import (
	"fmt"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/builder"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/product"
)

const (
	suvCarSeatsNumber       = 6
	suvCarEnginePower       = 1000
	suvCarTripComputerModel = "SUV car computer"
	suvCarGpsModel          = "Super GPS"

	sportCarSeatsNumber       = 1
	sportCarEnginePower       = 700
	sportCarTripComputerModel = "Super sports car computer"
	sportCarGpsModel          = "Cheap GPS"
)

var config = builder.Config{
	SuvCarParameters: &builder.CarParameters{
		SeatsNumber:       suvCarSeatsNumber,
		EnginePower:       suvCarEnginePower,
		TripComputerModel: suvCarTripComputerModel,
		GpsModel:          suvCarGpsModel,
	},
	SportCarParameters: &builder.CarParameters{
		SeatsNumber:       sportCarSeatsNumber,
		EnginePower:       sportCarEnginePower,
		TripComputerModel: sportCarTripComputerModel,
		GpsModel:          sportCarGpsModel,
	},
}

func main() {
	// строим машины
	carBuilder := builder.NewCarBuilder(product.NewCar)
	director := builder.NewCarDirector()
	director.SetBuilder(carBuilder)
	director.ConstructSuvCar(config.SuvCarParameters)
	suvCar := carBuilder.GetResult()
	director.ConstructSportsCar(config.SportCarParameters)
	sportsCar := carBuilder.GetResult()

	// пишем инструкции
	manualBuilder := builder.NewManualBuilder(product.NewManual)
	director.SetBuilder(manualBuilder)
	director.ConstructSuvCar(config.SuvCarParameters)
	suvCarManual := manualBuilder.GetResult()
	director.ConstructSportsCar(config.SportCarParameters)
	sportsCarManual := manualBuilder.GetResult()

	fmt.Println(suvCar)
	fmt.Println(sportsCar)
	fmt.Println(suvCarManual)
	fmt.Println(sportsCarManual)
}
