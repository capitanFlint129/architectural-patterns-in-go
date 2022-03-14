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

var (
	correctRequestTestCaseName = "Correct request"

	logger = logrus.New()
)

func Test_CookOrder(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputData
		expectedResult expectedResult
	}{
		{
			testCaseName: correctRequestTestCaseName,
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
			supportOperator := mocks.NewHandler()
			supportOperator.On("Handle", testData.inputData.request).Return(testData.expectedResult.solution, testData.expectedResult.error)
			chain := []handler{supportOperator}
			support := NewSupport(chain, logger)

			solution, err := support.ProcessRequest(testData.inputData.request)
			assert.Equal(t, solution, testData.expectedResult.solution)
			assert.ErrorIs(t, err, testData.expectedResult.error)

			supportOperator.AssertCalled(t, "Handle", testData.inputData.request)
			supportOperator.AssertNumberOfCalls(t, "Handle", testData.expectedResult.times)
		})
	}
}
