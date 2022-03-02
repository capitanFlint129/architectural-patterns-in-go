package builder

// Приватный интерфейс определяет контракт взаимодействия с этим пакетом
type builder interface {
	SetSeats(seatsNumber int)
	SetEngine(enginePower int)
	SetTripComputer(tripComputerModel string)
	SetGps(gpsModel string)
}

type CarDirector interface {
	SetBuilder(builder builder)
	ConstructSuvCar(seatsNumber int, enginePower int, tripComputerModel string, gpsModel string)
	ConstructSportsCar(seatsNumber int, enginePower int, tripComputerModel string, gpsModel string)
}

type carDirector struct {
	builder builder
}

func (c *carDirector) SetBuilder(builder builder) {
	c.builder = builder
}

// В данном случае матоды создания для разных типов автомобилей одинаковы,
// но так как в общем случае это может быть не так,то методы разделены
// для демонстрации паттерна
func (c *carDirector) ConstructSuvCar(seatsNumber int, enginePower int, tripComputerModel string, gpsModel string) {
	c.builder.SetSeats(seatsNumber)
	c.builder.SetEngine(enginePower)
	c.builder.SetTripComputer(tripComputerModel)
	c.builder.SetGps(gpsModel)
}

func (c *carDirector) ConstructSportsCar(seatsNumber int, enginePower int, tripComputerModel string, gpsModel string) {
	c.builder.SetSeats(seatsNumber)
	c.builder.SetEngine(enginePower)
	c.builder.SetTripComputer(tripComputerModel)
	c.builder.SetGps(gpsModel)
}

func NewCarDirector() CarDirector {
	return &carDirector{}
}
