package product

type Manual struct {
	seatsNumber       int
	enginePower       int
	tripComputerModel string
	gpsModel          string
}

func NewManual(seatsNumber int, enginePower int, tripComputerModel string, gpsModel string) Product {
	return &Manual{
		seatsNumber,
		enginePower,
		tripComputerModel,
		gpsModel,
	}
}
