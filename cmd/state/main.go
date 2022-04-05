package main

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/state/state"
	"github.com/sirupsen/logrus"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/state/ticket"
)

func main() {
	logger := logrus.New()
	newTicket := ticket.NewTicket(logger)
	initialState := state.NewDraftState(newTicket, logger)
	newTicket.SetState(initialState)
	// draft -> ready
	newTicket.Publish()
	// ready -> in progress
	newTicket.Complete()
	// in progress -> done
	newTicket.Complete()
	// done -> deleted
	newTicket.Delete()
}
