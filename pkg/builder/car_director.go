package builder

// Приватный интерфейс определяет контракт взаимодействия с этим пакетом
type builder interface {
	SetSeats(seatsNumber int)
	SetEngine(enginePower int)
	SetTripComputer(tripComputerModel string)
	SetGps(gpsModel string)
}

//CarDirector - director for cars creating
type CarDirector interface {
	SetBuilder(builder builder)
	ConstructSuvCar(parameters *CarParameters)
	ConstructSportsCar(parameters *CarParameters)
}

type carDirector struct {
	builder builder
}

//SetBuilder sets builder for director
func (c *carDirector) SetBuilder(builder builder) {
	c.builder = builder
}

// В данном случае матоды создания для разных типов автомобилей одинаковы,
// но так как в общем случае это может быть не так,то методы разделены
// для демонстрации паттерна

// ConstructSuvCar - manage builder to create suv car
func (c *carDirector) ConstructSuvCar(parameters *CarParameters) {
	c.builder.SetSeats(parameters.SeatsNumber)
	c.builder.SetEngine(parameters.EnginePower)
	c.builder.SetTripComputer(parameters.TripComputerModel)
	c.builder.SetGps(parameters.GpsModel)
}

// ConstructSportsCar - manage builder to create sport car
func (c *carDirector) ConstructSportsCar(parameters *CarParameters) {
	c.builder.SetSeats(parameters.SeatsNumber)
	c.builder.SetEngine(parameters.EnginePower)
	c.builder.SetTripComputer(parameters.TripComputerModel)
	c.builder.SetGps(parameters.GpsModel)
}

// NewCarDirector creates CarDirector
func NewCarDirector() CarDirector {
	return &carDirector{}
}
