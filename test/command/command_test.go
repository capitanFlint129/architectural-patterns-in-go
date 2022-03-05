package command

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/command"
	"github.com/capitanFlint129/architectural-patterns-in-go/test/command/mocks"
)

const wrongOrderErrorFmt = "No %s \n"

var (
	okOrdersData = []string{
		"Big Mac",
	}
	wrongOrdersData = []string{
		"Shaurma",
		"Pizza",
		"Sushi",
		"Hleb",
	}
)

func TestMakeOrder(t *testing.T) {
	for _, orderData := range okOrdersData {
		restaurantMock := mocks.NewRestaurant()
		makeOrder := command.NewMakeOrder(restaurantMock, orderData)
		restaurantMock.On("CookOrder", orderData).Return(nil)

		err := makeOrder.Execute()

		assert.Nil(t, err)
		restaurantMock.EXPECT().CookOrder(orderData).Return(nil).Times(1)
	}
}

func TestMakeWrongOrder(t *testing.T) {
	for _, orderData := range wrongOrdersData {
		restaurantMock := mocks.NewRestaurant()
		makeOrder := command.NewMakeOrder(restaurantMock, orderData)
		wrongOrderError := fmt.Errorf(wrongOrderErrorFmt, orderData)
		restaurantMock.On("CookOrder", orderData).Return(wrongOrderError)

		err := makeOrder.Execute()

		assert.Errorf(t, err, wrongOrderErrorFmt, orderData)
		restaurantMock.EXPECT().CookOrder(orderData).Return(nil).Times(1)
	}
}

func TestRequestMenu(t *testing.T) {
	restaurantMock := mocks.NewRestaurant()
	requestMenu := command.NewRequestMenu(restaurantMock)
	restaurantMock.On("GiveMenu").Return(nil)

	err := requestMenu.Execute()

	assert.Nil(t, err)
	restaurantMock.EXPECT().GiveMenu().Return(nil).Times(1)
}
