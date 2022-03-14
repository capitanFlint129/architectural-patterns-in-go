package handler

import "github.com/sirupsen/logrus"

// Handler is main interface for handlers
type Handler = interface {
	Handle(problem string, logger *logrus.Logger) (string, error)
}
