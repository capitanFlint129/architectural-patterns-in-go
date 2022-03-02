package product

type Manual struct {
	seatsNumber       int
	enginePower       int
	tripComputerModel string
	gpsModel          string
}

func ManualCreator(seatsNumber int, enginePower int, tripComputerModel string, gpsModel string) Product {
	return &Car{
		seatsNumber,
		enginePower,
		tripComputerModel,
		gpsModel,
	}
}
