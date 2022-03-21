package product

import "github.com/sirupsen/logrus"

type car struct {
	model  string
	speed  int
	logger *logrus.Logger
}

// Forward commands car to move forward
func (c *car) Forward() {
	c.logger.Infof("%s moves froward with speed %d", c.model, c.speed)
}

// Back commands car to move back
func (c *car) Back() {
	c.logger.Infof("%s moves back with speed %d", c.model, c.speed)
}

// GetModel returns car model
func (c *car) GetModel() string {
	return c.model
}

// GetSpeed returns car speed
func (c *car) GetSpeed() int {
	return c.speed
}

// NewCar creates new car
func NewCar(model string, speed int, logger *logrus.Logger) Product {
	return &car{
		model:  model,
		speed:  speed,
		logger: logger,
	}
}
