package state

import "github.com/sirupsen/logrus"

type inProgressState struct {
	ticket ticket
	logger *logrus.Logger
}

func (i *inProgressState) Publish() {
	i.logger.Info("inProgress: ticket already published")
}

func (i *inProgressState) Complete() {
	i.logger.Info("inProgress: complete")
	i.ticket.SetState(NewDoneState(i.ticket, i.logger))
}

func (i *inProgressState) Delete() {
	i.logger.Info("inProgress: delete")
	i.ticket.SetState(NewDeletedState(i.ticket, i.logger))
}

func NewInProgressState(ticket ticket, logger *logrus.Logger) State {
	return &inProgressState{
		ticket: ticket,
		logger: logger,
	}
}
