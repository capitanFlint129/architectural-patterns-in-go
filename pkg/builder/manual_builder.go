package builder

type manualBuilder struct {
	seatsNumber       int
	enginePower       int
	tripComputerModel string
	gpsModel          string
	manualCreator     CarCreator
}

func (m *manualBuilder) GetResult() product {
	return m.manualCreator(
		m.seatsNumber,
		m.enginePower,
		m.tripComputerModel,
		m.gpsModel,
	)
}

func (m *manualBuilder) SetSeats(seatsNumber int) {
	m.seatsNumber = seatsNumber
}

func (m *manualBuilder) SetEngine(enginePower int) {
	m.enginePower = enginePower
}

func (m *manualBuilder) SetTripComputer(tripComputerModel string) {
	m.tripComputerModel = tripComputerModel
}

func (m *manualBuilder) SetGps(gpsModel string) {
	m.gpsModel = gpsModel
}

func NewManualBuilder(manualCreator CarCreator) Builder {
	return &manualBuilder{
		manualCreator: manualCreator,
	}
}
