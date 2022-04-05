package strategy

import (
	"os"

	"github.com/sirupsen/logrus"
)

type docxStrategy struct {
	logger *logrus.Logger
}

func (d *docxStrategy) Convert(file *os.File) {
	d.logger.Info("docxStrategy: convert file")
}

func NewDocxStrategy(logger *logrus.Logger) Strategy {
	return &docxStrategy{
		logger: logger,
	}
}
