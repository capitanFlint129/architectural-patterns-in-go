package receiver

import "fmt"

// Receiver receives commands
type Receiver interface {
	GiveMenu()
	CookOrder(orderData string)
}

type restaurant struct{}

// GiveMenu provides menu to customer
func (r *restaurant) GiveMenu() {
	fmt.Println("Receiver: restaurant gives menu to customer")
}

// CookOrder cooks customers order
func (r *restaurant) CookOrder(orderData string) {
	fmt.Printf("Receiver: the chef prepares %s \n", orderData)
}

// NewRestaurant creates new commands receiver - restaurant
func NewRestaurant() Receiver {
	return &restaurant{}
}
