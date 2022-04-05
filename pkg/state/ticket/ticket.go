package ticket

import "github.com/sirupsen/logrus"

type state = interface {
	Publish()
	Complete()
	Delete()
}

type Ticket = interface {
	SetState(state state)
	Publish()
	Complete()
	Delete()
}

type ticket struct {
	state  state
	logger *logrus.Logger
}

func (t *ticket) SetState(state state) {
	t.state = state
}

func (t *ticket) Publish() {
	t.logger.Info("ticket: publish")
	t.state.Publish()
}

func (t *ticket) Complete() {
	t.logger.Info("ticket: complete")
	t.state.Complete()
}

func (t *ticket) Delete() {
	t.logger.Info("ticket: delete")
	t.state.Delete()
}

func NewTicket(
	logger *logrus.Logger,
) Ticket {
	return &ticket{
		logger: logger,
	}
}
