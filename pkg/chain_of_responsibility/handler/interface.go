package handler

// Handler is main interface for handlers
type Handler = interface {
	Handle(problem string) (string, error)
}
