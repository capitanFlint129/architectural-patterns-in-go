package builder

import "architectural-patterns-in-go/pkg/product"

type manualBuilderInterface interface {
	GetResult() product.Product

	setSeats(seatsNumber int)
	setEngine(enginePower int)
	setTripComputer(tripComputerModel string)
	setGps(gpsModel string)
}

type manualBuilder struct {
	seatsNumber       int
	enginePower       int
	tripComputerModel string
	gpsModel          string
	manualCreator     CarCreator
}

func (m *manualBuilder) GetResult() product.Product {
	return m.manualCreator(
		m.seatsNumber,
		m.enginePower,
		m.tripComputerModel,
		m.gpsModel,
	)
}

func (m *manualBuilder) setSeats(seatsNumber int) {
	m.seatsNumber = seatsNumber
}

func (m *manualBuilder) setEngine(enginePower int) {
	m.enginePower = enginePower
}

func (m *manualBuilder) setTripComputer(tripComputerModel string) {
	m.tripComputerModel = tripComputerModel
}

func (m *manualBuilder) setGps(gpsModel string) {
	m.gpsModel = gpsModel
}

func NewManualBuilder(manualCreator CarCreator) Builder {
	return &manualBuilder{
		manualCreator: manualCreator,
	}
}
