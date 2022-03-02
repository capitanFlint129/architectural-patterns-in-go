package product

type car struct {
	seatsNumber       int
	enginePower       int
	tripComputerModel string
	gpsModel          string
}

func NewCar(seatsNumber int, enginePower int, tripComputerModel string, gpsModel string) Product {
	return &car{
		seatsNumber:       seatsNumber,
		enginePower:       enginePower,
		tripComputerModel: tripComputerModel,
		gpsModel:          gpsModel,
	}
}
