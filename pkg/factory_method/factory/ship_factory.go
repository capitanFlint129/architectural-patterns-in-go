package factory

import "github.com/sirupsen/logrus"

type shipFactory struct {
	carCreator func(model string, speed int, logger *logrus.Logger) product
	logger     *logrus.Logger
}

// CreateProduct creates new ship
func (f *shipFactory) CreateProduct(model string, speed int) product {
	f.logger.Infof("Factory creates %s", model)
	return f.carCreator(model, speed, f.logger)
}

// NewShipFactory creates new ship factory
func NewShipFactory(productCreator productCreator, logger *logrus.Logger) Factory {
	return &shipFactory{
		carCreator: productCreator,
		logger:     logger,
	}
}
