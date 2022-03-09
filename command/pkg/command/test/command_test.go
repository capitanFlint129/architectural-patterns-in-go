package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/capitanFlint129/architectural-patterns-in-go/command/pkg/command"
	"github.com/capitanFlint129/architectural-patterns-in-go/command/pkg/command/test/mocks"
)

type inputData struct {
	restaurantName string
	restaurantMenu map[string]bool
	orderedDish    string
}

type expectedResult struct {
	error error
	times int
}

func Test_MakeOrder(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputData
		expectedResult expectedResult
	}{
		{
			testCaseName: "Dish in restaurant menu",
			inputData: inputData{
				orderedDish: "Big Mac",
			},
			expectedResult: expectedResult{
				error: nil,
				times: 1,
			},
		},
		{
			testCaseName: "Dish not in menu",
			inputData: inputData{
				orderedDish: "Shaurma",
			},
			expectedResult: expectedResult{
				error: fmt.Errorf("No Shaurma \n"),
				times: 1,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			restaurantMock := mocks.NewRestaurant()
			makeOrder := command.NewMakeOrder(restaurantMock, testData.inputData.orderedDish)
			restaurantMock.On("CookOrder", testData.inputData.orderedDish).Return(testData.expectedResult.error)

			err := makeOrder.Execute()
			if err == nil {
				assert.ErrorIs(t, err, testData.expectedResult.error)
			} else {
				assert.EqualError(t, err, testData.expectedResult.error.Error())
			}
			restaurantMock.EXPECT().CookOrder(testData.inputData.orderedDish).Return(testData.expectedResult.error).Times(testData.expectedResult.times)
		})
	}
}

func Test_RequestMenu(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputData
		expectedResult expectedResult
	}{
		{
			testCaseName: "Dish in restaurant menu",
			inputData: inputData{
				orderedDish: "Big Mac",
			},
			expectedResult: expectedResult{
				error: nil,
				times: 1,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			restaurantMock := mocks.NewRestaurant()
			requestMenu := command.NewRequestMenu(restaurantMock)
			restaurantMock.On("GiveMenu").Return(testData.expectedResult.error)

			err := requestMenu.Execute()

			if err == nil {
				assert.ErrorIs(t, err, testData.expectedResult.error)
			} else {
				assert.EqualError(t, err, testData.expectedResult.error.Error())
			}
			restaurantMock.EXPECT().GiveMenu().Return(testData.expectedResult.error).Times(testData.expectedResult.times)
		})
	}
}
