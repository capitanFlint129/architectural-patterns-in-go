package product

import "github.com/sirupsen/logrus"

type car struct {
	Model  string
	Speed  int
	logger *logrus.Logger
}

// Forward commands car to move forward
func (c *car) Forward() {
	c.logger.Infof("%s moves froward with Speed %d", c.Model, c.Speed)
}

// Back commands car to move back
func (c *car) Back() {
	c.logger.Infof("%s moves back with Speed %d", c.Model, c.Speed)
}

// NewCar creates new car
func NewCar(model string, speed int, logger *logrus.Logger) Product {
	return &car{
		Model:  model,
		Speed:  speed,
		logger: logger,
	}
}
