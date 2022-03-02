package builder

// определяет контракт взаимодействия с этим пакетом
type builder interface {
	setSeats(seatsNumber int)
	setEngine(enginePower int)
	setTripComputer(tripComputerModel string)
	setGps(gpsModel string)
}

type CarDirector interface {
	SetBuilder(builder builder)
	ConstructSuvCar()
	ConstructSportsCar()
}

type carDirector struct {
	builder builder
}

func (c *carDirector) SetBuilder(builder builder) {
	c.builder = builder
}

func (c *carDirector) ConstructSuvCar() {
	c.builder.setSeats(6)
	c.builder.setEngine(1000)
	c.builder.setTripComputer("SUV car computer")
	c.builder.setGps("Super GPS")
}

func (c *carDirector) ConstructSportsCar() {
	c.builder.setSeats(1)
	c.builder.setEngine(700)
	c.builder.setTripComputer("Super sports car computer")
	c.builder.setGps("Cheap GPS")
}

func NewCarDirector() CarDirector {
	return &carDirector{}
}
