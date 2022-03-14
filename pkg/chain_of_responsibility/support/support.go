package support

import (
	"github.com/sirupsen/logrus"
)

type handler = interface {
	Handle(problem string, logger *logrus.Logger) (string, error)
}

// Support helps user with some troubles
type Support interface {
	ProcessRequest(request string, logger *logrus.Logger) (string, error)
}

type support struct {
	chain []handler
}

// ProcessRequest initiate request processing
func (s *support) ProcessRequest(request string, logger *logrus.Logger) (string, error) {
	logger.Info("Support start request processing")
	return s.chain[0].Handle(request, logger)
}

// NewSupport creates new support
func NewSupport(chain []handler) Support {
	return &support{chain: chain}
}
