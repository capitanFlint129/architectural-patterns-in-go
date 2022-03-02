package builder

import "github.com/capitanFlint129/architectural-patterns-in-go/pkg/product"

type carBuilder struct {
	seatsNumber       int
	enginePower       int
	tripComputerModel string
	gpsModel          string
	carCreator        CarCreator // functor - скрываем конструктор за параметром и не привязываемся к конкретному
}

func (c *carBuilder) GetResult() product.Product {
	return c.carCreator(
		c.seatsNumber,
		c.enginePower,
		c.tripComputerModel,
		c.gpsModel,
	)
}

func (c *carBuilder) setSeats(seatsNumber int) {
	c.seatsNumber = seatsNumber
}

func (c *carBuilder) setEngine(enginePower int) {
	c.enginePower = enginePower
}

func (c *carBuilder) setTripComputer(tripComputerModel string) {
	c.tripComputerModel = tripComputerModel
}

func (c *carBuilder) setGps(gpsModel string) {
	c.gpsModel = gpsModel
}

func NewCarBuilder(carCreator CarCreator) Builder {
	return &carBuilder{
		carCreator: carCreator,
	}
}
