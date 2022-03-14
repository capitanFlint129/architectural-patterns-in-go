package support

import (
	"github.com/sirupsen/logrus"
)

type handler = interface {
	Handle(problem string) (string, error)
}

// Support helps user with some troubles
type Support interface {
	ProcessRequest(request string) (string, error)
}

type support struct {
	chain  []handler
	logger *logrus.Logger
}

// ProcessRequest initiate request processing
func (s *support) ProcessRequest(request string) (string, error) {
	s.logger.Info("Support start request processing")
	return s.chain[0].Handle(request)
}

// NewSupport creates new support
func NewSupport(chain []handler, logger *logrus.Logger) Support {
	return &support{
		chain:  chain,
		logger: logger,
	}
}
