package factory

import "github.com/sirupsen/logrus"

type carFactory struct {
	carCreator func(model string, speed int, logger *logrus.Logger) product
	logger     *logrus.Logger
}

// CreateProduct creates new car
func (f *carFactory) CreateProduct(model string, speed int) product {
	f.logger.Infof("carFactory creates %s", model)
	return f.carCreator(model, speed, f.logger)
}

// NewCarFactory creates new car factory
func NewCarFactory(productCreator productCreator, logger *logrus.Logger) Factory {
	return &carFactory{
		carCreator: productCreator,
		logger:     logger,
	}
}
