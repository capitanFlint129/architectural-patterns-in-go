package receiver

import (
	"testing"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/receiver"
	"github.com/stretchr/testify/assert"
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
		restaurant := receiver.NewRestaurant()
		err := restaurant.CookOrder(orderData)

		assert.Nil(t, err)
	}
}

func TestMakeWrongOrder(t *testing.T) {
	for _, orderData := range wrongOrdersData {
		restaurant := receiver.NewRestaurant()
		err := restaurant.CookOrder(orderData)

		assert.Errorf(t, err, wrongOrderErrorFmt, orderData)
	}
}

func TestRequestMenu(t *testing.T) {
	restaurant := receiver.NewRestaurant()
	err := restaurant.GiveMenu()

	assert.Nil(t, err)
}
