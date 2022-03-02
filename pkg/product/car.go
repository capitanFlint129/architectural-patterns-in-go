package product

type Car struct {
	seatsNumber       int
	enginePower       int
	tripComputerModel string
	gpsModel          string
}

func CarCreator(seatsNumber int, enginePower int, tripComputerModel string, gpsModel string) Product {
	return &Car{
		seatsNumber,
		enginePower,
		tripComputerModel,
		gpsModel,
	}
}
