package command

import "fmt"

// TODO глянуть логирование (zap, gokit, logras)
// TODO глянуть фреймворки (gokit, fasthttp, nethttp, gorilla/mux)

type makeOrder struct {
	restaurant restaurant
	dish       string
}

// Execute - sends the visitor's order
func (m *makeOrder) Execute() error {
	fmt.Println("Command: makeOrder executes")
	err := m.restaurant.CookOrder(m.dish)
	return err
}

// NewMakeOrder creates new makeOrder command
func NewMakeOrder(restaurant restaurant, dish string) Command {
	return &makeOrder{
		restaurant: restaurant,
		dish:       dish,
	}
}