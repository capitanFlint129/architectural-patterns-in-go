package builder

type manualBuilder struct {
	seatsNumber       int
	enginePower       int
	tripComputerModel string
	gpsModel          string
	manualCreator     CarCreator
}

// GetResult returns created manual for car
func (m *manualBuilder) GetResult() product {
	return m.manualCreator(
		m.seatsNumber,
		m.enginePower,
		m.tripComputerModel,
		m.gpsModel,
	)
}

// SetSeats defines seats described in car manual
func (m *manualBuilder) SetSeats(seatsNumber int) {
	m.seatsNumber = seatsNumber
}

// SetEngine defines engine described in car manual
func (m *manualBuilder) SetEngine(enginePower int) {
	m.enginePower = enginePower
}

// SetTripComputer defines trip computer described in car manual
func (m *manualBuilder) SetTripComputer(tripComputerModel string) {
	m.tripComputerModel = tripComputerModel
}

// SetGps defines gps module described in car manual
func (m *manualBuilder) SetGps(gpsModel string) {
	m.gpsModel = gpsModel
}

//NewManualBuilder creates new manual for car
func NewManualBuilder(manualCreator CarCreator) Builder {
	return &manualBuilder{
		manualCreator: manualCreator,
	}
}
