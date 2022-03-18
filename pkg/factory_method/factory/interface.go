package factory

import "github.com/sirupsen/logrus"

type product = interface {
	Forward()
	Back()
}

type productCreator func(model string, speed int, logger *logrus.Logger) product

// Factory is common factory interface
type Factory interface {
	// CreateProduct returns new product created by factory
	CreateProduct(model string, speed int) product
}
