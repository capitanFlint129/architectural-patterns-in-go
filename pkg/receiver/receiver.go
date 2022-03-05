package receiver

import "fmt"

type Receiver interface {
	GiveMenu()
	CookOrder(orderData string)
}

type restaurant struct{}

func (r *restaurant) GiveMenu() {
	fmt.Println("Receiver: restaurant gives menu to customer")
}

func (r *restaurant) CookOrder(orderData string) {
	fmt.Printf("Receiver: the chef prepares %s \n", orderData)
}

func NewRestaurant() Receiver {
	return &restaurant{}
}
