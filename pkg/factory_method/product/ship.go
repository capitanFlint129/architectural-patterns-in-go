package product

import "github.com/sirupsen/logrus"

type ship struct {
	Model  string
	Speed  int
	logger *logrus.Logger
}

// Forward commands ship to move forward
func (s *ship) Forward() {
	s.logger.Infof("%s moves froward with Speed %d", s.Model, s.Speed)
}

// Back commands ship to move back
func (s *ship) Back() {
	s.logger.Infof("%s moves back with Speed %d", s.Model, s.Speed)
}

// NewShip creates new ship
func NewShip(model string, speed int, logger *logrus.Logger) Product {
	return &ship{
		Model:  model,
		Speed:  speed,
		logger: logger,
	}
}
