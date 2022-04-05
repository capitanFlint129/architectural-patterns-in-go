package pdf_converter

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/strategy/pdf_converter/mocks"
)

const (
	convertTestCaseName = "Convert file via pdf converter"
)

type inputData struct {
}

type expectedResult struct {
	numberOfCallsOfStrategyConvertMethod int
}

var logger = logrus.New()

func Test_PdfConverter(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputData
		expectedResult expectedResult
	}{
		{
			testCaseName: convertTestCaseName,
			expectedResult: expectedResult{
				numberOfCallsOfStrategyConvertMethod: 1,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			strategyMock := mocks.NewStrategyMock()
			strategyMock.On("Convert", mock.AnythingOfType("*os.File")).Return()

			pdfConverter := NewPdfConverter(logger)
			pdfConverter.SetStrategy(strategyMock)
			var file *os.File
			pdfConverter.Convert(file)

			strategyMock.AssertNumberOfCalls(t, "Convert", testData.expectedResult.numberOfCallsOfStrategyConvertMethod)
		})
	}
}
