package main

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/factory_method/factory"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/factory_method/product"
	"github.com/sirupsen/logrus"
)

const (
	sportCarModel = "Sport Car"
	sportCarSpeed = 500

	boatModel = "Boat"
	boatSpeed = 100
)

func main() {
	logger := logrus.New()

	carFactory := factory.NewCarFactory(product.NewCar, logger)
	sportCar := carFactory.CreateProduct(sportCarModel, sportCarSpeed)
	sportCar.Forward()
	sportCar.Back()

	shipFactory := factory.NewShipFactory(product.NewShip, logger)
	ship := shipFactory.CreateProduct(boatModel, boatSpeed)
	ship.Forward()
	ship.Back()
}
