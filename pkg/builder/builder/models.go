package builder

// Config define parameters for different car types
type Config struct {
	SuvCarParameters   *CarParameters
	SportCarParameters *CarParameters
}

// CarParameters defines parameters for a specific car type
type CarParameters struct {
	SeatsNumber       int
	EnginePower       int
	TripComputerModel string
	GpsModel          string
}
