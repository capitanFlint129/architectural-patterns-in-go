package state

import "github.com/sirupsen/logrus"

type doneState struct {
	ticket ticket
	logger *logrus.Logger
}

func (d *doneState) Publish() {
	d.logger.Info("done: ticket already published")
}

func (d *doneState) Complete() {
	d.logger.Info("done: ticket already done")
}

func (d *doneState) Delete() {
	d.logger.Info("done: delete")
	d.ticket.SetState(NewDeletedState(d.ticket, d.logger))
}

func NewDoneState(ticket ticket, logger *logrus.Logger) State {
	return &doneState{
		ticket: ticket,
		logger: logger,
	}
}
