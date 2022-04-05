package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/strategy/pdf_converter"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/strategy/strategy"
)

const fileType = "docx"

func main() {
	logger := logrus.New()

	converter := pdf_converter.NewPdfConverter(logger)

	docxStrategy := strategy.NewDocxStrategy(logger)
	pptxStrategy := strategy.NewPptxStrategy(logger)
	jpgStrategy := strategy.NewJpgStrategy(logger)

	var file *os.File

	switch fileType {
	case "docx":
		converter.SetStrategy(docxStrategy)
	case "pptx":
		converter.SetStrategy(pptxStrategy)
	case "jpg":
		converter.SetStrategy(jpgStrategy)
	}
	converter.Convert(file)
}
