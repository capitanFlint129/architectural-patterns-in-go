package builder

type carBuilder struct {
	seatsNumber       int
	enginePower       int
	tripComputerModel string
	gpsModel          string
	carCreator        CarCreator // functor - скрываем конструктор за параметром и не привязываемся к конкретному
}

func (c *carBuilder) GetResult() product {
	return c.carCreator(
		c.seatsNumber,
		c.enginePower,
		c.tripComputerModel,
		c.gpsModel,
	)
}

func (c *carBuilder) SetSeats(seatsNumber int) {
	c.seatsNumber = seatsNumber
}

func (c *carBuilder) SetEngine(enginePower int) {
	c.enginePower = enginePower
}

func (c *carBuilder) SetTripComputer(tripComputerModel string) {
	c.tripComputerModel = tripComputerModel
}

func (c *carBuilder) SetGps(gpsModel string) {
	c.gpsModel = gpsModel
}

func NewCarBuilder(carCreator CarCreator) Builder {
	return &carBuilder{
		carCreator: carCreator,
	}
}
