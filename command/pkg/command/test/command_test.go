package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/capitanFlint129/architectural-patterns-in-go/command/pkg/command"
	"github.com/capitanFlint129/architectural-patterns-in-go/command/pkg/command/test/mocks"
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

// TODO добавить инициализацию в начало теста
// TODO запускать разные сценарии теста в t.run()
// TODO использовать структуру для табличного теста: (начальные данные, результат, доп инфа)
func Test_MakeOrder(t *testing.T) {
	for _, dish := range okOrdersData {
		restaurantMock := mocks.NewRestaurant()
		makeOrder := command.NewMakeOrder(restaurantMock, dish)
		restaurantMock.On("CookOrder", dish).Return(nil)

		err := makeOrder.Execute()
		assert.Nil(t, err)
		restaurantMock.EXPECT().CookOrder(dish).Return(nil).Times(1)
	}
}

func TestMakeWrongOrder(t *testing.T) {
	for _, dish := range wrongOrdersData {
		restaurantMock := mocks.NewRestaurant()
		makeOrder := command.NewMakeOrder(restaurantMock, dish)
		wrongOrderError := fmt.Errorf(wrongOrderErrorFmt, dish)
		restaurantMock.On("CookOrder", dish).Return(wrongOrderError)

		err := makeOrder.Execute()
		assert.Errorf(t, err, wrongOrderErrorFmt, dish)
		restaurantMock.EXPECT().CookOrder(dish).Return(nil).Times(1)
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
