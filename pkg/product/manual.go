package product

type manual struct {
	seatsNumber       int
	enginePower       int
	tripComputerModel string
	gpsModel          string
}

func NewManual(seatsNumber int, enginePower int, tripComputerModel string, gpsModel string) Product {
	return &manual{
		seatsNumber:       seatsNumber,
		enginePower:       enginePower,
		tripComputerModel: tripComputerModel,
		gpsModel:          gpsModel,
	}
}
