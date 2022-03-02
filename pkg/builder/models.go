package builder

type Config struct {
	SuvCarParameters   CarParameters
	SportCarParameters CarParameters
}

type CarParameters struct {
	SeatsNumber       int
	EnginePower       int
	TripComputerModel string
	GpsModel          string
}
