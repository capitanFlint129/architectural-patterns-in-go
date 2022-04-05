package state

type ticket = interface {
	SetState(state State)
	Publish()
	Complete()
	Delete()
}

type State = interface {
	Publish()
	Complete()
	Delete()
}
