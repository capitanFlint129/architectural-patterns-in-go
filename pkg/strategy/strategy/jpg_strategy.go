package strategy

import (
	"os"

	"github.com/sirupsen/logrus"
)

type jpgStrategy struct {
	logger *logrus.Logger
}

func (d *jpgStrategy) Convert(file *os.File) {
	d.logger.Info("jpgStrategy: convert file")
}

func NewJpgStrategy(logger *logrus.Logger) Strategy {
	return &jpgStrategy{
		logger: logger,
	}
}
