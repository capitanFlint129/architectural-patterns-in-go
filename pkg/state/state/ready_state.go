package state

import "github.com/sirupsen/logrus"

type readyState struct {
	ticket ticket
	logger *logrus.Logger
}

func (r *readyState) Publish() {
	r.logger.Info("ready: ticket already published")
}

func (r *readyState) Complete() {
	r.logger.Info("ready: complete")
	r.ticket.SetState(NewInProgressState(r.ticket, r.logger))
}

func (r *readyState) Delete() {
	r.logger.Info("ready: delete")
	r.ticket.SetState(NewDeletedState(r.ticket, r.logger))
}

func NewReadyState(ticket ticket, logger *logrus.Logger) State {
	return &readyState{
		ticket: ticket,
		logger: logger,
	}
}
