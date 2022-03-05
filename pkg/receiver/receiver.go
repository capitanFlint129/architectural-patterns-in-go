package receiver

import "fmt"

// Receiver receives commands
type Receiver interface {
	GiveMenu() error
	CookOrder(orderData string) error
}

type restaurant struct{}

// GiveMenu provides menu to customer
func (r *restaurant) GiveMenu() error {
	fmt.Println("Receiver: restaurant gives menu to customer")
	return nil
}

// CookOrder cooks customers order
func (r *restaurant) CookOrder(orderData string) error {
	fmt.Printf("Receiver: the chef prepares %s \n", orderData)
	if orderData != "Big Mac" {
		return fmt.Errorf("No %s \n", orderData)
	}
	return nil
}

// NewRestaurant creates new commands receiver - restaurant
func NewRestaurant() Receiver {
	return &restaurant{}
}
