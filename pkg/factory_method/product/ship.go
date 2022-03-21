package product

import "github.com/sirupsen/logrus"

type ship struct {
	model  string
	speed  int
	logger *logrus.Logger
}

// Forward commands ship to move forward
func (s *ship) Forward() {
	s.logger.Infof("%s moves froward with speed %d", s.model, s.speed)
}

// Back commands ship to move back
func (s *ship) Back() {
	s.logger.Infof("%s moves back with speed %d", s.model, s.speed)
}

// GetModel returns product model
func (s *ship) GetModel() string {
	return s.model
}

// GetSpeed returns product speed
func (s *ship) GetSpeed() int {
	return s.speed
}

// NewShip creates new ship
func NewShip(model string, speed int, logger *logrus.Logger) Product {
	return &ship{
		model:  model,
		speed:  speed,
		logger: logger,
	}
}
