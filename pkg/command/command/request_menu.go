package command

import (
	"github.com/sirupsen/logrus"
)

type requestMenu struct {
	restaurant restaurant
}

// Execute - requests menu for customer
func (r *requestMenu) Execute() error {
	logrus.Info("Command: requestMenu executes")
	err := r.restaurant.GiveMenu()
	return err
}

// NewRequestMenu creates new requestMenu command
func NewRequestMenu(restaurant restaurant) Command {
	return &requestMenu{restaurant: restaurant}
}
