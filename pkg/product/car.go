package product

type Car struct {
	seatsNumber       int
	enginePower       int
	tripComputerModel string
	gpsModel          string
}

func NewCar(seatsNumber int, enginePower int, tripComputerModel string, gpsModel string) Product {
	return &Car{
		seatsNumber,
		enginePower,
		tripComputerModel,
		gpsModel,
	}
}
