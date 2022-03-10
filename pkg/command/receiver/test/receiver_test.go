package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/command/receiver"
)

type inputData struct {
	restaurantName string
	restaurantMenu map[string]bool
	orderedDish    string
}

type expectedResult struct {
	error error
}

func Test_MakeOrder(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputData
		expectedResult expectedResult
	}{
		{
			testCaseName: "Dish in menu",
			inputData: inputData{
				restaurantName: "mcdonalds",
				restaurantMenu: map[string]bool{
					"Big Mac": true,
				},
				orderedDish: "Big Mac",
			},
			expectedResult: expectedResult{
				error: nil,
			},
		},
		{
			testCaseName: "Dish not in menu",
			inputData: inputData{
				restaurantName: "mcdonalds",
				restaurantMenu: map[string]bool{
					"Big Mac": true,
				},
				orderedDish: "Shaurma",
			},
			expectedResult: expectedResult{
				error: fmt.Errorf("No Shaurma \n"),
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			restaurant := receiver.NewRestaurant(testData.inputData.restaurantName, testData.inputData.restaurantMenu)
			err := restaurant.CookOrder(testData.inputData.orderedDish)
			if err == nil {
				assert.ErrorIs(t, err, testData.expectedResult.error)
			} else {
				assert.EqualError(t, err, testData.expectedResult.error.Error())
			}
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
			testCaseName: "Restaurant gives menu",
			inputData: inputData{
				restaurantName: "mcdonalds",
				restaurantMenu: map[string]bool{
					"Big Mac": true,
				},
			},
			expectedResult: expectedResult{
				error: nil,
			},
		},
	} {
		restaurant := receiver.NewRestaurant(testData.inputData.restaurantName, testData.inputData.restaurantMenu)
		err := restaurant.GiveMenu()
		if err == nil {
			assert.ErrorIs(t, err, testData.expectedResult.error)
		} else {
			assert.EqualError(t, err, testData.expectedResult.error.Error())
		}
	}
}
