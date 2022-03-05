package command

import "fmt"

type requestMenu struct {
	restaurant restaurant
}

// Execute - requests menu for customer
func (r *requestMenu) Execute() {
	fmt.Println("Command: requestMenu executes")
	r.restaurant.GiveMenu()
}

// NewrequestMenu creates new requestMenu command
func NewRequestMenu(restaurant restaurant) Command {
	return &requestMenu{restaurant: restaurant}
}
