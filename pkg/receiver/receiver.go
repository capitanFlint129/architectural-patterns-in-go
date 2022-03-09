package receiver

import (
	"fmt"
	"time"
)

// Receiver receives commands
type Receiver = interface {
	GiveMenu() error
	CookOrder(dish string) error
}

type restaurant struct {
	name string
	menu map[string]bool
}

// GiveMenu provides menu to customer
func (r *restaurant) GiveMenu() error {
	fmt.Printf("%s: restaurant gives menu to customer\n", r.name)
	time.Sleep(time.Second)
	fmt.Println("Receiver: restaurant gives menu to customer")
	return nil
}

// CookOrder cooks customers order
func (r *restaurant) CookOrder(dish string) error {
	fmt.Printf("Receiver: the chef prepares %s \n", dish)
	if _, ok := r.menu[dish]; ok == false {
		return fmt.Errorf("No %s \n", dish)
	}
	return nil
}

// NewRestaurant creates new commands receiver - restaurant
func NewRestaurant(name string, menu map[string]bool) Receiver {
	return &restaurant{
		name: name,
		menu: menu,
	}
}
