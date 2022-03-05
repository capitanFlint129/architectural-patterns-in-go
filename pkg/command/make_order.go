package command

import "fmt"

type makeOrder struct {
	restaurant restaurant
	orderData  string
}

// Execute - sends the visitor's order
func (m *makeOrder) Execute() {
	fmt.Println("Command: makeOrder executes")
	m.restaurant.CookOrder(m.orderData)
}

// NewMakeOrder creates new makeOrder command
func NewMakeOrder(restaurant restaurant, orderData string) Command {
	return &makeOrder{
		restaurant: restaurant,
		orderData:  orderData,
	}
}