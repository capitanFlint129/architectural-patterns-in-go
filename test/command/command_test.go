package command

import (
	"testing"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/command"
	"github.com/capitanFlint129/architectural-patterns-in-go/test/command/mocks"
)

const orderData = "Big Mac"

func TestMakeOrder(t *testing.T) {
	restaurantMock := mocks.NewRestaurant()
	makeOrder := command.NewMakeOrder(restaurantMock, orderData)
	restaurantMock.On("CookOrder", orderData).Return()

	makeOrder.Execute()

	restaurantMock.EXPECT().CookOrder(orderData).Return().Times(1)
}

func TestRequestMenu(t *testing.T) {
	restaurantMock := mocks.NewRestaurant()
	requestMenu := command.NewRequestMenu(restaurantMock)
	restaurantMock.On("GiveMenu").Return()

	requestMenu.Execute()

	restaurantMock.EXPECT().GiveMenu().Return().Times(1)
}
