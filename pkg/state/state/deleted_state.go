package state

import "github.com/sirupsen/logrus"

type deletedState struct {
	ticket ticket
	logger *logrus.Logger
}

func (d *deletedState) Publish() {
	d.logger.Info("deleted: can't publish deleted ticket")
}

func (d *deletedState) Complete() {
	d.logger.Info("deleted: can't complete deleted ticket")
}

func (d *deletedState) Delete() {
	d.logger.Info("deleted: ticket already deleted")
}

func NewDeletedState(ticket ticket, logger *logrus.Logger) State {
	return &deletedState{
		ticket: ticket,
		logger: logger,
	}
}
