package factory

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/factory_method/factory/mocks"
)

type shipData struct {
	model string
	speed int
}

type inputDataShipFactory struct {
	createdShip shipData
}

var (
	shipFactoryTestCaseName = "Successful ship creation"

	loggerShipFactory = logrus.New()
)

func Test_ShipFactory(t *testing.T) {
	for _, testData := range []struct {
		testCaseName string
		inputData    inputDataShipFactory
	}{
		{
			testCaseName: shipFactoryTestCaseName,
			inputData: inputDataShipFactory{
				createdShip: shipData{
					model: "ship",
					speed: 100,
				},
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			ship := mocks.NewProduct()
			shipCreator := func(model string, speed int, logger *logrus.Logger) product {
				return ship
			}
			shipFactory := NewShipFactory(shipCreator, loggerShipFactory)
			shipFormFactory := shipFactory.CreateProduct(testData.inputData.createdShip.model, testData.inputData.createdShip.speed)

			assert.Equal(t, ship, shipFormFactory)
		})
	}
}
