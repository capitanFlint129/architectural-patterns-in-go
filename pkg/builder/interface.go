package builder

type product = interface{}

// Builder have all methods to produce some car product
type Builder interface {
	GetResult() product
	SetSeats(seatsNumber int)
	SetEngine(enginePower int)
	SetTripComputer(tripComputerModel string)
	SetGps(gpsModel string)
}

// CarCreator creates new car
type CarCreator func(seatsNumber int, enginePower int, tripComputerModel string, gpsModel string) product
