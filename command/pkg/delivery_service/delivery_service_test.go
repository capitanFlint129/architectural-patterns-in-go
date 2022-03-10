package delivery_service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/capitanFlint129/architectural-patterns-in-go/command/pkg/delivery_service/mocks"
)

type inputData struct {
	restaurantMapMocked           map[string]restaurant
	requestMenuCommandMockCreator func(restaurant restaurant) command
	cookOrderCommandMockCreator   func(restaurant restaurant, dish string) command
	orderedDish                   string
}

type expectedResult struct {
	error error
	times int
}

func Test_CookOrder(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputData
		expectedResult expectedResult
	}{
		{
			testCaseName: "Correct order",
			inputData: inputData{
				restaurantMapMocked: map[string]restaurant{
					"mcdonalds": mocks.NewRestaurant(),
				},
				orderedDish: "Big Mac",
			},
			expectedResult: expectedResult{
				error: nil,
				times: 1,
			},
		},
		{
			testCaseName: "Wrong order",
			inputData: inputData{
				restaurantMapMocked: map[string]restaurant{
					"mcdonalds": mocks.NewRestaurant(),
				},
				orderedDish: "Big Mac",
			},
			expectedResult: expectedResult{
				error: fmt.Errorf("No Big Mac \n"),
				times: 1,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			cookOrderCommandMock := mocks.NewCommand()
			cookOrderCommandMock.On("Execute").Return(testData.expectedResult.error)
			deliveryService := NewDeliveryService(
				testData.inputData.restaurantMapMocked,
				func(restaurant restaurant) command {
					return mocks.NewCommand()
				},
				func(restaurant restaurant, dish string) command {
					return cookOrderCommandMock
				},
			)

			err := deliveryService.MakeOrder("mcdonalds", testData.inputData.orderedDish)
			assert.ErrorIs(t, err, testData.expectedResult.error)
			cookOrderCommandMock.EXPECT().Execute().Return(testData.expectedResult.error).Times(testData.expectedResult.times)
		})
	}
}

func Test_RequestMenus(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputData
		expectedResult expectedResult
	}{
		{
			testCaseName: "Request menus",
			inputData: inputData{
				restaurantMapMocked: map[string]restaurant{
					"mcdonalds": mocks.NewRestaurant(),
				},
			},
			expectedResult: expectedResult{
				error: nil,
				times: 1,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			requestMenuCommandMock := mocks.NewCommand()
			requestMenuCommandMock.On("Execute").Return(testData.expectedResult.error)
			delisveryService := NewDeliveryService(
				testData.inputData.restaurantMapMocked,
				func(restaurant restaurant) command {
					return requestMenuCommandMock
				},
				func(restaurant restaurant, dish string) command {
					return mocks.NewCommand()
				},
			)

			err := delisveryService.RequestMenus()
			assert.ErrorIs(t, err, testData.expectedResult.error)
			requestMenuCommandMock.EXPECT().Execute().Return(testData.expectedResult.error).Times(testData.expectedResult.times)
		})
	}
}
