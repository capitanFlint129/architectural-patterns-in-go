package builder

type carBuilder struct {
	seatsNumber       int
	enginePower       int
	tripComputerModel string
	gpsModel          string
	carCreator        CarCreator // functor - скрываем конструктор за параметром и не привязываемся к конкретному
}

// GetResult returns created car
func (c *carBuilder) GetResult() product {
	return c.carCreator(
		c.seatsNumber,
		c.enginePower,
		c.tripComputerModel,
		c.gpsModel,
	)
}

// SetSeats sets seats for car
func (c *carBuilder) SetSeats(seatsNumber int) {
	c.seatsNumber = seatsNumber
}

// SetEngine sets engine for car
func (c *carBuilder) SetEngine(enginePower int) {
	c.enginePower = enginePower
}

// SetTripComputer sets trip computer for car
func (c *carBuilder) SetTripComputer(tripComputerModel string) {
	c.tripComputerModel = tripComputerModel
}

// SetGps sets gps module for car
func (c *carBuilder) SetGps(gpsModel string) {
	c.gpsModel = gpsModel
}

//NewCarBuilder creates new car
func NewCarBuilder(carCreator CarCreator) Builder {
	return &carBuilder{
		carCreator: carCreator,
	}
}
