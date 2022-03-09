package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/capitanFlint129/architectural-patterns-in-go/command/pkg/receiver"
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
	for _, dish := range okOrdersData {
		restaurant := receiver.NewRestaurant()
		err := restaurant.CookOrder(dish)

		assert.Nil(t, err)
	}
}

func TestMakeWrongOrder(t *testing.T) {
	for _, dish := range wrongOrdersData {
		restaurant := receiver.NewRestaurant()
		err := restaurant.CookOrder(dish)

		assert.Errorf(t, err, wrongOrderErrorFmt, dish)
	}
}

func TestRequestMenu(t *testing.T) {
	restaurant := receiver.NewRestaurant()
	err := restaurant.GiveMenu()

	assert.Nil(t, err)
}
