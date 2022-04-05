package pdf_converter

import (
	"github.com/sirupsen/logrus"
	"os"
)

type strategy interface {
	Convert(file *os.File)
}

type PdfConverter interface {
	SetStrategy(strategy strategy)
	Convert(file *os.File)
}

type pdfConverter struct {
	strategy strategy
	logger   *logrus.Logger
}

func (p *pdfConverter) SetStrategy(strategy strategy) {
	p.logger.Info("pdfConverter: set strategy")
	p.strategy = strategy
}

func (p *pdfConverter) Convert(file *os.File) {
	p.logger.Info("pdfConverter: use strategy")
	p.strategy.Convert(file)
}

func NewPdfConverter(logger *logrus.Logger) PdfConverter {
	return &pdfConverter{
		logger: logger,
	}
}
