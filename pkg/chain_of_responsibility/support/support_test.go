package support

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/chain_of_responsibility/support/mocks"
)

type inputData struct {
	request string
}

type expectedResult struct {
	solution string
	error    error
	times    int
}

func Test_CookOrder(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputData
		expectedResult expectedResult
	}{
		{
			testCaseName: "Correct request",
			inputData: inputData{
				request: "laptop broken",
			},
			expectedResult: expectedResult{
				solution: "laptop broken",
				error:    nil,
				times:    1,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			logger := logrus.New()
			supportOperator := mocks.NewHandler()
			supportOperator.On("Handle", testData.inputData.request, logger).Return(testData.expectedResult.solution, testData.expectedResult.error)
			chain := []handler{supportOperator}

			support := NewSupport(chain)
			solution, err := support.ProcessRequest(testData.inputData.request, logger)

			assert.Equal(t, solution, testData.expectedResult.solution)
			assert.ErrorIs(t, err, testData.expectedResult.error)
			supportOperator.AssertCalled(t, "Handle", testData.inputData.request, logger)
			supportOperator.AssertNumberOfCalls(t, "Handle", testData.expectedResult.times)
		})
	}
}
