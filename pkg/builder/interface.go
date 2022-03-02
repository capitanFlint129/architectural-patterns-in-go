package builder

type product = interface{}

type Builder interface {
	GetResult() product

	SetSeats(seatsNumber int)
	SetEngine(enginePower int)
	SetTripComputer(tripComputerModel string)
	SetGps(gpsModel string)
}

type CarCreator func(seatsNumber int, enginePower int, tripComputerModel string, gpsModel string) product
