package strategy

import (
	"os"

	"github.com/sirupsen/logrus"
)

type pptxStrategy struct {
	logger *logrus.Logger
}

func (d *pptxStrategy) Convert(file *os.File) {
	d.logger.Info("pptxStrategy: convert file")
}

func NewPptxStrategy(logger *logrus.Logger) Strategy {
	return &pptxStrategy{
		logger: logger,
	}
}
