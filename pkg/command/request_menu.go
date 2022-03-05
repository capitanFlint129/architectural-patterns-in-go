package command

import "fmt"

type requestMenu struct {
	restaurant restaurant
}

// Execute
func (r *requestMenu) Execute() {
	fmt.Println("Command: requestMenu executes")
	r.restaurant.GiveMenu()
}

// NewrequestMenu
func NewRequestMenu(restaurant restaurant) Command {
	return &requestMenu{restaurant: restaurant}
}
