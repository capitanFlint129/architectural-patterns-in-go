package product

type car struct {
	seatsNumber       int
	enginePower       int
	tripComputerModel string
	gpsModel          string
}

// NewCar creates new car with specified parameters
func NewCar(seatsNumber int, enginePower int, tripComputerModel string, gpsModel string) Product {
	return &car{
		seatsNumber:       seatsNumber,
		enginePower:       enginePower,
		tripComputerModel: tripComputerModel,
		gpsModel:          gpsModel,
	}
}
