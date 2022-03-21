package factory

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/factory_method/factory/mocks"
)

type carData struct {
	model string
	speed int
}

type inputDataCarFactory struct {
	createdCar carData
}

var (
	carFactoryTestCaseName = "Successful car creation"

	loggerCarFactory = logrus.New()
)

func Test_CarFactory(t *testing.T) {
	for _, testData := range []struct {
		testCaseName string
		inputData    inputDataCarFactory
	}{
		{
			testCaseName: carFactoryTestCaseName,
			inputData: inputDataCarFactory{
				createdCar: carData{
					model: "car",
					speed: 100,
				},
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			car := mocks.NewProduct()
			carCreator := func(model string, speed int, logger *logrus.Logger) product {
				return car
			}
			carFactory := NewCarFactory(carCreator, loggerCarFactory)
			carFormFactory := carFactory.CreateProduct(testData.inputData.createdCar.model, testData.inputData.createdCar.speed)

			assert.Equal(t, car, carFormFactory)
		})
	}
}
