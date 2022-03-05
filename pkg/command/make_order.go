package command

import "fmt"

type makeOrder struct {
	restaurant restaurant
	orderData  string
}

// Execute
func (m *makeOrder) Execute() {
	fmt.Println("Command: makeOrder executes")
	m.restaurant.CookOrder(m.orderData)
}

// NewMakeOrder
func NewMakeOrder(restaurant restaurant, orderData string) Command {
	return &makeOrder{
		restaurant: restaurant,
		orderData:  orderData,
	}
}
