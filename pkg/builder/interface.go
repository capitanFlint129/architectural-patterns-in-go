package builder

import "architectural-patterns-in-go/pkg/product"

type Builder interface {
	GetResult() product.Product

	setSeats(seatsNumber int)
	setEngine(enginePower int)
	setTripComputer(tripComputerModel string)
	setGps(gpsModel string)
}

type CarParameters struct {
	SeatsNumber       int
	EnginePower       int
	TripComputerModel string
	GpsModel          string
}

type CarCreator func(seatsNumber int, enginePower int, tripComputerModel string, gpsModel string) product.Product
