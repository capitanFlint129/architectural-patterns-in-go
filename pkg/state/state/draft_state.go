package state

import "github.com/sirupsen/logrus"

type draftState struct {
	ticket ticket
	logger *logrus.Logger
}

func (d *draftState) Publish() {
	d.logger.Info("draft: publish")
	d.ticket.SetState(NewReadyState(d.ticket, d.logger))
}

func (d *draftState) Complete() {
	d.logger.Info("draft: can't complete draft")
}

func (d *draftState) Delete() {
	d.logger.Info("draft: delete")
	d.ticket.SetState(NewDeletedState(d.ticket, d.logger))
}

func NewDraftState(ticket ticket, logger *logrus.Logger) State {
	return &draftState{
		ticket: ticket,
		logger: logger,
	}
}
